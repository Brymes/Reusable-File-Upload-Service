package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)


var MongoClient *mongo.Database

func InitDb() {
	uri := os.Getenv("DATABASE_URI")

	if uri == "" {
		//uri = "mongodb://127.0.0.1:27017"
		uri = "mongodb://host.docker.internal:27017"
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Println("Error Connecting to Mongo")
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Println("Error Connecting to Mongo")
		log.Fatal(err)
	}

	//TODO Change DB Name
	MongoClient = client.Database("Sabre-E-Certificates")
}


