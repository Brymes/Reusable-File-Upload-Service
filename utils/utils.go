package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"runtime/debug"
)

type APIResponse map[string]interface{}

func GenerateUniqueID(length int) (string, error) {
	bytes := make([]byte, length)

	chars := "0123456789abcdefghijklmnopqrstuvwxyz"

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}

func ConvertTagsToString(tags []string) (tagString string) {
	// Tags Parsing for S3 Integration
	for _, element := range tags {
		tagString += fmt.Sprintf("%v=&", element)
	}
	return
}

func HandlePanic(response APIResponse, logger *log.Logger) {
	if err := recover(); err != nil {
		logger.Println(string(debug.Stack()))
		response["status"] = false
		response["message"] = err
	}
}
