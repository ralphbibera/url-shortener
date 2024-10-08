package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func RedirectUrlHandler(lambdaContext context.Context, lambdaInput events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/*
		checks if the short url id exists in the databse
		if exists{
			return redirect url
		} else{
		 	return error
		}
	*/
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "ok",
	}, nil
}
