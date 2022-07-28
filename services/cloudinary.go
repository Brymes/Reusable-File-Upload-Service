package services

import (
	"Upload-Service/config"
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"log"
	"os"
)

type Cloudinary struct{}

func (c Cloudinary) Upload(s ServicePayload, logger *log.Logger) (url, publicID string) {
	ctx := context.Background()

	// Use CallBack or notification_url param to perform async uploads
	resp, err := config.CloudinaryInstance.Client.Upload.Upload(ctx, s.Filepath, uploader.UploadParams{Tags: s.Tags})
	if err != nil {
		logger.Panicf("Error Uploading File: %v", err)
	}

	url, publicID = resp.SecureURL, resp.PublicID

	return
}

func (c Cloudinary) SignUpload(s ServicePayload, notificationURL string, logger *log.Logger) string {
	//NOTE
	// Append parameters to the form data. The parameters that are signed using
	// the signing function (signuploadform) need to match these.
	paramsToSign, err := api.StructToParams(uploader.UploadParams{Tags: s.Tags, NotificationURL: notificationURL})
	if err != nil {
		logger.Panicf("Bad Request: %v", err)
	}

	resp, err := api.SignParameters(paramsToSign, os.Getenv("CLOUDINARY_API_SECRET"))
	if err != nil {
		logger.Panicf("Error Signing URLL %v", err)
	}

	logger.Println(resp)

	return fmt.Sprintf("https://api.cloudinary.com/v1_1/%v/auto/upload", resp)

}
