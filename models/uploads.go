package models

import (
	"Upload-Service/config"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Upload struct {
	ID primitive.ObjectID `bson:"_id,  omitempty" json:"-"`
	// Implement
	Identifier string   `bson:"postID" json:"postID" binding:"required"`
	UploadURL  string   `bson:"upload_url" json:"upload_url"`
	Tags       []string `bson:"tags" json:"tags"`
}

func (u *Upload) SaveUpload(logger *log.Logger) {
	//TODO Implement uniqueness for  Field Identifier

	ctx := context.TODO()
	_, err := config.MongoClient.Collection("Uploads").InsertOne(ctx, u)
	if err != nil {
		logger.Panic("Error saving upload")
	}

	return
}

func (u *Upload) DeleteUpload() {

}

func (u *Upload) FetchUpload() {

}
