package entities

import (
	"log"
	"testing"
	"url-shortener/pkg/logger"
	"url-shortener/pkg/rds"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
}

func TestGenerateUrlId(t *testing.T) {
	url := Url{}
	url.GenerateUrlId()
	logger.Log(url)
}

func TestGetItem(t *testing.T) {
	database, err := rds.Connect()
	if err != nil {
		t.Error(err)
	}
	defer database.Close()

	url := Url{
		Database:    database,
		UrlId:       aws.String("Rb57P1eW"),
		RedirectUrl: aws.String("www.google.com"),
	}

	err = url.GetUrl()
	if err != nil {
		t.Error(err)
	}

	logger.Log(url)
}

func TestDoesUrlExist(t *testing.T) {
	database, err := rds.Connect()
	if err != nil {
		t.Error(err)
	}
	defer database.Close()

	url := Url{
		Database:    database,
		UrlId:       aws.String("Rb57PeW"),
		RedirectUrl: aws.String("www.google.com"),
	}

	exist, err := url.DoesUrlExist()
	if err != nil {
		t.Error(err)
	}

	logger.Log(exist)
}
