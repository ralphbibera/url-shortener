package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func ShortenUrlHandler(lambdaContext context.Context, lambdaInput events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/*
		checks if the long url exists in the databse
		if exists{
			return existing shortened url
		} else{
		 	create shortened url item
			store to databsae
			return item
		 }
	*/

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "success",
	}, nil
}
