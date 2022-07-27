package api

import (
	"Upload-Service/config"
	"Upload-Service/services"
	u "Upload-Service/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

func UploadFile(c *gin.Context) {
	var request services.ServicePayload

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Cannot Parse Upload %s", err.Error())
		return
	}

	filename := fmt.Sprintf("uploads/%v", filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, "Upload File Error: %s", err.Error())
		return
	}

	request.Filepath, request.Tags, request.Service = filename, u.ParseTagsFromFormData(c.PostFormArray("tags")), c.PostForm("service")

	reqBuffer, reqLogger := config.InitRequestLogger(request.Service)

	url, publicID := request.UploadFile(reqLogger)

	log.Println(reqBuffer)
	c.IndentedJSON(http.StatusOK, gin.H{
		"message":  "File Upload Successful",
		"url":      url,
		"publicID": publicID,
	})
}

func GetUploadURL(c *gin.Context) {
	var request services.ServicePayload

	serverURL := "https://" + c.Request.Host + c.Request.URL.String()

	if err := c.BindJSON(&request); err != nil {
		//TODO return error message
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	reqBuffer, reqLogger := config.InitRequestLogger(request.Service)

	response := request.CreateUploadURL(serverURL, reqLogger)

	log.Println(reqBuffer)

	if response["status"].(bool) != true {
		c.IndentedJSON(http.StatusBadRequest, response)
	}
	c.IndentedJSON(http.StatusOK, response)
}
