package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"strings"
)

func LogErr(err error, logger *log.Logger) {
	if err != nil {
		logger.Panic(err)
	}
}

func GenerateUniqueID(length int) (string, error) {
	bytes := make([]byte, length)

	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}

func ConvertTagsToString(tags map[string]interface{}) (tagString string) {
	// Tags Parsing for S3 Integration
	for key, element := range tags {
		tagString += fmt.Sprintf("%v=%v&", key, element)
	}
	return
}

func ParseTagsFromFormData(tagArray []string) (tagMap map[string]interface{}) {
	for _, element := range tagArray {
		chars := strings.Split(element, "=")
		tagMap[chars[0]] = chars[1]
	}
	return
}
