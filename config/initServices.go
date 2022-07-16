package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"log"
)

var CloudinaryInstance *cloudinary.Cloudinary

func InitCloudinary() {
	var err error

	CloudinaryInstance, err = cloudinary.New()

	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}
}
