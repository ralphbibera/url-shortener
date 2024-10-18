package wrapper

import (
	"context"
	"url-shortener/pkg/logger"

	"github.com/aws/aws-lambda-go/lambda"
)

type lambdaFunction[Input any, Output any] func(context context.Context, lambdaInput Input) (Output, error)

func Start[Input any, Output any](lambdaFunction lambdaFunction[Input, Output]) {
	lambda.Start(func(context context.Context, lambdaInput Input) (Output, error) {
		logger.Log(lambdaInput, "Lambda Input:")
		response, err := lambdaFunction(context, lambdaInput)
		logger.Log(response, "Lambda Response:")
		return response, err
	})
}
