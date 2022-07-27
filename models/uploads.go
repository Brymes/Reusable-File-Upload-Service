package models

import (
	"Upload-Service/config"
	"Upload-Service/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Upload struct {
	ID primitive.ObjectID `bson:"_id,  omitempty" json:"-"`
	// Implement
	Identifier string                 `bson:"postID" json:"postID" binding:"required"`
	UploadURL  string                 `bson:"upload_url" json:"upload_url"`
	Tags       map[string]interface{} `bson:"tags" json:"tags"`
}

func (u Upload) SaveUpload(logger *log.Logger) {
	//TODO Implement uniqueness for  Field Identifier

	ctx := context.TODO()
	_, err := config.MongoClient.Collection("Uploads").InsertOne(ctx, u)
	utils.LogErr(err, logger)

	return
}

func (u Upload) DeleteUpload() {

}
