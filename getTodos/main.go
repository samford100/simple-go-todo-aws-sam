package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Todo struct {
	Id   string `json:"id"`
	Desc string `json:"desc"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get from dynamo
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		//ExpressionAttributeNames:  expr.Names(),
		//ExpressionAttributeValues: expr.Values(),
		//FilterExpression:          expr.Filter(),
		//ProjectionExpression:      expr.Projection(),
		TableName: aws.String("todos"),
	}

	result, err := svc.Scan(params)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if *result.Count == 0 {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 404,
		}, nil
	}

	var todos []Todo
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &todos)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	// End dynamo

	// Send response back to API gateway
	response, err := json.Marshal(todos)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
