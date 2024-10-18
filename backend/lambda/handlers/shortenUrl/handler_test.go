package main

import (
	"context"
	"log"
	"testing"
	"url-shortener/pkg/logger"

	"github.com/aws/aws-lambda-go/events"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
}

func TestShortenUrlHandler(t *testing.T) {
	response, err := ShortenUrlHandler(context.Background(), events.APIGatewayProxyRequest{
		Body: "brave.com",
	})
	if err != nil {
		t.Error(err)
	}
	logger.Log(response)
}
