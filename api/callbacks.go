package api

import (
	"Upload-Service/config"
	"Upload-Service/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CloudinaryNotification(c *gin.Context) {

	var notificationPayload models.Upload
	//TODO to use map[string]interface or not

	if err := c.BindJSON(&notificationPayload); err != nil {
		//TODO return error message
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	reqBuffer, reqLogger := config.InitRequestLogger("cloudinary")
	defer log.Println(reqBuffer)

	notificationPayload.SaveUpload(reqLogger)

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error processing Upload from Cloudinary")
			log.Println(r)
			log.Println(notificationPayload)
		}
	}()
	c.IndentedJSON(http.StatusOK, "")
}
