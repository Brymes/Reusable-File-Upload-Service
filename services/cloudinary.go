package services

import (
	"Upload-Service/config"
	"Upload-Service/utils"
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
	utils.LogErr(err, logger)

	url, publicID = resp.SecureURL, resp.PublicID

	return
}

func (c Cloudinary) SignUpload(s ServicePayload, notificationURL string, logger *log.Logger) string {
	//NOTE
	// Append parameters to the form data. The parameters that are signed using
	// the signing function (signuploadform) need to match these.
	paramsToSign, err := api.StructToParams(uploader.UploadParams{Tags: s.Tags, NotificationURL: notificationURL})
	utils.LogErr(err, logger)

	resp, err := api.SignParameters(paramsToSign, os.Getenv("CLOUDINARY_API_SECRET"))
	utils.LogErr(err, logger)

	logger.Println(resp)

	return fmt.Sprintf("https://api.cloudinary.com/v1_1/%v/auto/upload", resp)

}
