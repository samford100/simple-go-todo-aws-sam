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

	id := request.QueryStringParameters["id"]
	//return events.APIGatewayProxyResponse{
	//	Body:       id,
	//	StatusCode: 200,
	//}, nil

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("todos"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if result.Item == nil {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 404,
		}, nil
	}

	todo := Todo{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &todo)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	// End dynamo

	// Send response back to API gateway
	response, err := json.Marshal(todo)

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
