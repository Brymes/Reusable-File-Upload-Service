package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudinary/cloudinary-go/v2"
	"log"
	"os"
)

var (
	CloudinaryInstance *CloudinaryClient
	DefaultBucket      string
	S3Instance         *S3Client
)

type S3Client struct {
	Client   *s3.Client
	Bucket   string
	IsActive bool
}

type CloudinaryClient struct {
	Client   *cloudinary.Cloudinary
	IsActive bool
}

func InitCloudinary() {
	var err error
	CloudinaryInstance.IsActive = true

	if os.Getenv("CLOUDINARY_API_SECRET") == "" {
		log.Println("Failed to initialize Cloudinary" + err.Error())
		CloudinaryInstance.IsActive = false
	}

	CloudinaryInstance.Client, err = cloudinary.New()

	if err != nil {
		log.Println("Failed to initialize Cloudinary" + err.Error())
		CloudinaryInstance.IsActive = false
	}
}

func InitS3() {
	S3Instance.IsActive = true

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Println("configuration error, " + err.Error())
		S3Instance.IsActive = false
	}

	S3Instance.Client = s3.NewFromConfig(cfg)
	err = InitBucketS3()

	if err != nil {
		S3Instance.IsActive = false
	}

}

func InitBucketS3() error {
	bucket := os.Getenv("S3_DEFAULT_BUCKET")
	if bucket == "" {
		bucket = "Uploads"
	}

	input := &s3.CreateBucketInput{
		Bucket: &bucket,
	}

	_, err := S3Instance.Client.CreateBucket(context.TODO(), input)

	return err
}
