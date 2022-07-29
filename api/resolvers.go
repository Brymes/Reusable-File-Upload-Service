package api

import (
	"Upload-Service/config"
	"Upload-Service/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

//TODO Use middlewares to validate if service exists and is active or not

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

	request.Filepath, request.Tags, request.Service = filename, c.PostFormArray("tags"), c.PostForm("service")

	reqBuffer, reqLogger := config.InitRequestLogger(request.Service)

	response := request.UploadFile(reqLogger)

	defer log.Println(reqBuffer)

	if response["status"].(bool) != true {
		c.IndentedJSON(http.StatusInternalServerError, response)
	} else {
		c.IndentedJSON(http.StatusOK, response)
	}

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

	defer log.Println(reqBuffer)

	if response["status"].(bool) != true {
		c.IndentedJSON(http.StatusBadRequest, response)
	} else {
		c.IndentedJSON(http.StatusOK, response)
	}
}
