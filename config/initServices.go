package config

import (
	"Upload-Service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/cloudinary/cloudinary-go/v2"
	"log"
	"os"
)

var (
	CloudinaryInstance *CloudinaryClient
	S3Instance         *S3Client
)

type S3Client struct {
	Client        *s3.Client
	DefaultBucket string
	Region        string
	IsActive      bool
}

type CloudinaryClient struct {
	Client   *cloudinary.Cloudinary
	IsActive bool
}

func InitCloudinary() {
	var err error
	CI := CloudinaryClient{IsActive: true}

	if os.Getenv("CLOUDINARY_API_SECRET") == "" {
		log.Println("Failed to initialize Cloudinary")
		CI.IsActive = false
	}

	CI.Client, err = cloudinary.New()

	if err != nil {
		log.Println("Failed to initialize Cloudinary" + err.Error())
		CI.IsActive = false
	}
	CloudinaryInstance = &CI
}

func InitS3() {
	var (
		bucketExists *types.BucketAlreadyExists
		opErr        *smithy.OperationError
	)

	S3I := S3Client{IsActive: true}

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Println("configuration error, " + err.Error())
		S3I.IsActive = false
	}

	S3I.Client = s3.NewFromConfig(cfg)
	err = InitBucketS3(S3I)

	if err == nil {
		S3I.IsActive = true
		S3Instance = &S3I
	} else if errors.As(err, &bucketExists) {
		log.Printf("Bucket %s Already Exists:", S3I.DefaultBucket)
	} else if errors.As(err, &opErr) {
		log.Println("")
		log.Println("The Program encountered an Operation error whilst creating bucket, Perhaps you should specify another bucket name. See Rules at https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html")
		log.Println("")
		log.Printf("Failed to call service: %s, operation: %s, error: %v", opErr.Service(), opErr.Operation(), opErr.Unwrap())
	} else {
		log.Println("Error")
		log.Println(err)
		S3I.IsActive = false
	}
}

func InitBucketS3(S3I S3Client) error {
	S3I.DefaultBucket = os.Getenv("S3_DEFAULT_BUCKET")
	if S3I.DefaultBucket == "" {
		randNum, err := utils.GenerateUniqueID(10)
		if err != nil {
			log.Fatalf("Something Terrible happened %v", err)
		}
		S3I.DefaultBucket = fmt.Sprintf("brymes-upload-service-%v", randNum)
		log.Println("Auto Create Bucket:", S3I.DefaultBucket)
	}

	input := &s3.CreateBucketInput{
		Bucket: &S3I.DefaultBucket,
	}

	_, err := S3I.Client.CreateBucket(context.TODO(), input)

	return err
}
