package main

import "context"

func RedirectShortUrlHandler(lambdaContext context.Context, lambdaInput any) (any, error) {
	return "success", nil
}
