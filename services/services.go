package services

import (
	"Upload-Service/models"
	"Upload-Service/utils"
	"log"
)

type ServicePayload struct {
	Filepath string   `json:"-"`
	UploadID string   `json:"upload_id"`
	Service  string   `json:"service"`
	Bucket   string   `json:"bucket"`
	Tags     []string `json:"tags"`
}

func (s ServicePayload) UploadFile(logger *log.Logger) (response utils.APIResponse) {
	response = utils.APIResponse{"status": false}

	defer utils.HandlePanic(response, logger)

	service := AllServices[s.Service]

	if service == nil {
		response["url"] = ""
		response["publicID"] = ""
		panic("Unknown Service Selected")
	}

	url, publicID := service.Upload(s, logger)

	upload := models.Upload{
		Identifier: publicID,
		UploadURL:  url,
		Tags:       s.Tags,
	}
	upload.SaveUpload(logger)

	response["url"], response["publicID"] = url, publicID
	response["status"], response["message"] = true, "File Upload Successful"

	return

}

func (s ServicePayload) CreateUploadURL(serverURL string, logger *log.Logger) (response utils.APIResponse) {
	response = utils.APIResponse{"status": false}

	defer utils.HandlePanic(response, logger)

	service := AllServices[s.Service]

	serverURL = serverURL + "/" + s.Service

	if service == nil {
		response["signedURL"] = ""
		panic("Unknown Service Selected")
	}

	uploadURL := service.SignUpload(s, serverURL, logger)
	response["status"], response["message"], response["signedURL"] = true, "Upload URL generated successfully", uploadURL
	return
}

func (s ServicePayload) DeleteUpload() {
	//service := AllServices[s.Service]

}

type Service interface {
	Upload(s ServicePayload, logger *log.Logger) (url, publicID string)
	SignUpload(s ServicePayload, notificationURL string, logger *log.Logger) string
}

var AllServices = map[string]Service{
	"cloudinary": Cloudinary{},
	"s3":         S3{},
}
