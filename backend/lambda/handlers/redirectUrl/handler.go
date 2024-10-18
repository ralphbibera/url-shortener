package main

import (
	"context"
	"net/http"
	"url-shortener/pkg/api"
	"url-shortener/pkg/entities"
	"url-shortener/pkg/rds"

	"github.com/aws/aws-lambda-go/events"
)

func RedirectUrlHandler(lambdaContext context.Context, lambdaInput events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	database, err := rds.Connect()
	if err != nil {
		return api.InternalServerError(), err
	}
	defer database.Close()

	url := entities.Url{
		UrlId:    &lambdaInput.Body,
		Database: database,
	}

	exist, err := url.DoesUrlExist()
	if err != nil {
		return api.InternalServerError(), err
	}
	if exist {
		url.GetUrl()
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       http.StatusText(http.StatusNotFound),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       *url.RedirectUrl,
	}, nil
}
