package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

	id := request.PathParameters["id"]

	var todo Todo
	err := json.Unmarshal([]byte(request.Body), &todo)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	todo.Id = id

	// Will reject if ID does not exist?
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String(todo.Desc),
			},
		},
		TableName: aws.String("todos"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(todo.Id),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#mydesc": aws.String("desc"),
		},
		UpdateExpression: aws.String("set #mydesc = :r"),
	}

	_, err = svc.UpdateItem(input)

	if err != nil {
		b, _ := json.Marshal(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 503,
			Body:       string(b),
		}, err
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
