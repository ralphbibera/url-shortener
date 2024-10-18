package api

import (
	"testing"
	"url-shortener/pkg/logger"
)

func TestInternalServerError(t *testing.T) {
	response := InternalServerError()
	logger.Log(response)
}
