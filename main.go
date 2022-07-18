package main

import (
	"Upload-Service/api"
	"Upload-Service/config"
)

func init() {
	config.InitDb()
	config.InitCloudinary()
}

func main() {
	api.Server()
}
