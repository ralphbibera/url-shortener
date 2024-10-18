package api

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func InternalServerError() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}
}
