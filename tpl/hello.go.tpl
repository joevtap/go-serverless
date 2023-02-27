package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := `{"message": "Hello World!"}`
	return events.APIGatewayProxyResponse{Body: response, StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}