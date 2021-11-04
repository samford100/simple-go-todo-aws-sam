package main

//https://{restapi_id}.execute-api.{region}.amazonaws.com/{stage_name}/
//https://cjo5jmgvm3.execute-api.us-east-1.amazonaws.com/a58sdi/
//https://cjo5jmgvm3.execute-api.us-east-1.amazonaws.com/Stage

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
	// Add to dynamo
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	todo := Todo{
		Id:   "1",
		Desc: "New todo",
	}

	av, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("todos"),
	}

	_, err = svc.PutItem(input)

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
