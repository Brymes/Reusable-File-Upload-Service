package main

import (
	"Upload-Service/api"
	"Upload-Service/config"
)

func init() {
	config.InitDb()
	config.InitCloudinary()
	config.InitS3()
}

func main() {
	api.Server()
}
