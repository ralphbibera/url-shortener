package main

import (
	"url-shortener/pkg/wrapper"
)

func main() {
	wrapper.Start(ShortenUrlHandler)
}
