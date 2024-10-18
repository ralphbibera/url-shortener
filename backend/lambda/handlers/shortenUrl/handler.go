package main

import (
	"context"
	"net/http"
	"url-shortener/pkg/api"
	"url-shortener/pkg/entities"
	"url-shortener/pkg/rds"

	"github.com/aws/aws-lambda-go/events"
)

func ShortenUrlHandler(lambdaContext context.Context, lambdaInput events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	database, err := rds.Connect()
	if err != nil {
		return api.InternalServerError(), err
	}
	defer database.Close()

	url := entities.Url{
		RedirectUrl: &lambdaInput.Body,
		Database:    database,
	}

	exist, err := url.DoesUrlExist()
	if err != nil {
		return api.InternalServerError(), err
	}

	if exist {
		err := url.GetUrl()
		if err != nil {
			return api.InternalServerError(), err
		}
	} else {
		url.GenerateUrlId()
		url.Insert()
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       *url.UrlId,
	}, nil
}
