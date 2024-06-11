package utils

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionString = "mongodb://localhost:27017"
	dbName           = "userDatabase"
	collectionName   = "user"
)

var Collection *mongo.Collection

// create database connection
func init() {
	clientOptions := options.Client().ApplyURI(connectionString) // Connection to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("MongoDB Connection instance is ready")
}
