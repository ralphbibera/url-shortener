package logger

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func Log(input any, message ...string) {
	byteSlice, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	structuredLog := strings.ReplaceAll(string(byteSlice), "\n", "\r")
	if len(message) > 0 {
		log.Print(message[0]+" ", structuredLog)
	} else {
		log.Print(structuredLog)
	}
}

func Error(err error) {
	Log(err.Error(), "Error:")
}

func Fatal(err error) {
	Log(err.Error(), "Fatal Error:")
	os.Exit(1)
}
