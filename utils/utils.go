package utils

import (
	"crypto/rand"
	"log"
)

func LogErr(err error, logger *log.Logger) {
	if err != nil {
		logger.Panic(err)
	}
}

func GenerateUniqueID() (string, error) {
	bytes := make([]byte, 13)

	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}
