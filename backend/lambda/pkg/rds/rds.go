package rds

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"url-shortener/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	_DATA_SOURCE_NAME_FORMAT = "postgres://%s:%s@%s:%s"
	_DATABASE_SECRET_ARN     = "DATABASE_SECRET_ARN"
	_DATABASE_ADDRESS        = "DATABASE_ADDRESS"
	_DATABASE_PORT           = "DATABASE_PORT"
	_DATABASE_DRIVER         = "pgx"
)

type credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func Connect() (*sql.DB, error) {
	awsConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	secretsmanagerClient := secretsmanager.NewFromConfig(awsConfig)

	rdsSecret, err := secretsmanagerClient.GetSecretValue(context.Background(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(os.Getenv(_DATABASE_SECRET_ARN)),
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var databaseCredentials credentials
	err = json.Unmarshal([]byte(*rdsSecret.SecretString), &databaseCredentials)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	dataSourceName := fmt.Sprintf(
		_DATA_SOURCE_NAME_FORMAT,
		databaseCredentials.Username,
		url.QueryEscape(databaseCredentials.Password),
		os.Getenv(_DATABASE_ADDRESS),
		os.Getenv(_DATABASE_PORT),
	)

	database, err := sql.Open(_DATABASE_DRIVER, dataSourceName)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return database, nil
}
