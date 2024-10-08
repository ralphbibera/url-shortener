package logger

import (
	"encoding/json"
	"log"
	"strings"
)

func Log(message string, input any) {
	byteSlice, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	structuredLog := strings.ReplaceAll(string(byteSlice), "\n", "\r")
	log.Print(message+" ", structuredLog)
}
