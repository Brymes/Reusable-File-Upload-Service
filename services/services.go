package services

import (
	"Upload-Service/models"
	"github.com/gin-gonic/gin"
	"log"
)

type ServicePayload struct {
	Filepath string   `json:"-"`
	UploadID string   `json:"upload_id"`
	Service  string   `json:"service"`
	Tags     []string `json:"tags"`
}

func (s ServicePayload) UploadFile(logger *log.Logger) (url, publicID string) {
	service := AllServices[s.Service]

	url, publicID = service.Upload(s, logger)

	upload := models.Upload{
		Identifier: publicID,
		UploadURL:  url,
		Tags:       s.Tags,
	}

	upload.SaveUpload(logger)

	return

}

func (s ServicePayload) CreateUploadURL(serverURL string, logger *log.Logger) gin.H {
	service := AllServices[s.Service]

	serverURL = serverURL + "/" + s.Service

	if service == nil {
		return gin.H{
			"status":    false,
			"message":   "Unknown Service Selected",
			"signedURL": "",
		}
	}

	uploadURL := service.SignUpload(s, serverURL, logger)

	return gin.H{
		"status":    true,
		"message":   "Upload URL generated successfully",
		"signedURL": uploadURL,
	}
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
}
