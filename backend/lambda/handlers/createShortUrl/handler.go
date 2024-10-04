package main

import "context"

func CreateShortUrlHandler(lambdaContext context.Context, lambdaInput any) (any, error) {
	return "success", nil
}
