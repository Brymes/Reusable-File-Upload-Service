package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func Server() {

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	//config.AllowHeaders = []string{"Origin", "Authorization", "Content-Length", "Content-Type"}
	//config.ExposeHeaders = []string{"Content-Length"}
	//config.AllowOrigins = []string{"*"}

	router := gin.Default()
	router.Use(cors.New(config))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// TODO implement JWT verification
	/*	jwtMiddleware := JwtMiddleWareConfig()
			errInit := jwtMiddleware.MiddlewareInit()

			if errInit != nil {
				log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
			}

		router.POST("/login", jwtMiddleware.LoginHandler)

		auth := router.Group("/")

		auth.Use(jwtMiddleware.MiddlewareFunc())
		auth.POST("register", register)
	*/

	router.POST("upload", UploadFile)
	router.POST("signUpload", GetUploadURL)

	//Notification/CallBack Routes
	router.POST("cloudinary", CloudinaryNotification)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8888"
	}
	err := router.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
		return
	}
}
