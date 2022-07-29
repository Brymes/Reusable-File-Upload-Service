package services

import (
	"Upload-Service/config"
	"Upload-Service/utils"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct{}

func (s3Payload S3) Upload(s ServicePayload, logger *log.Logger) (url, publicID string) {
	if config.S3Instance.IsActive != true {
		logger.Panic("Service Not Activated, Kindly Set Auth Keys")
	}

	file, err := os.Open(s.Filepath)

	if err != nil {
		logger.Panicf("Unable to Open file %v for upload", s.Filepath)
	}

	defer file.Close()

	publicID, _ = utils.GenerateUniqueID(13)
	if publicID == "" {
		logger.Panic("Internal Server Error Contact Support")
	}

	tags := utils.ConvertTagsToString(s.Tags)
	input := &s3.PutObjectInput{
		Bucket:  &config.S3Instance.DefaultBucket,
		Key:     &publicID,
		Body:    file,
		Tagging: &tags,
	}

	if s.Bucket != "" {
		input.Bucket = &s.Bucket
	}
	_, err = config.S3Instance.Client.PutObject(context.TODO(), input)

	if err != nil {
		logger.Panicf("Internal Server Error: %v", err.Error())
	}
	url = ConstructObjectURL(publicID)

	return
}

func (s3Payload S3) SignUpload(s ServicePayload, notificationURL string, logger *log.Logger) string {
	if config.S3Instance.IsActive != true {
		logger.Panic("Service Not Activated, Kindly Set Auth Keys")
	}

	id, _ := utils.GenerateUniqueID(13)
	if id == "" {
		logger.Panic("Internal Server Error")
	}
	input := &s3.GetObjectInput{
		Bucket: &config.S3Instance.DefaultBucket,
		Key:    &id,
	}
	if s.Bucket != "" {
		input.Bucket = &s.Bucket
	}

	api := s3.NewPresignClient(config.S3Instance.Client)
	resp, err := api.PresignGetObject(context.TODO(), input)

	if err != nil || resp.URL == "" {
		logger.Panic("Internal Server Error whilst signing URL: %v", err)
	}

	return resp.URL
}

func ConstructObjectURL(objectKey string) (url string) {
	s3Instance := config.S3Instance
	url = fmt.Sprintf("https://%v.amazonaws.com/%v/%v", s3Instance.Region, s3Instance.DefaultBucket, objectKey)
	return
}
