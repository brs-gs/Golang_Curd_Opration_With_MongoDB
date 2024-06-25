package utils

import (
	"context"
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
	log.Println("Starting MongoDB connection")
	clientOptions := options.Client().ApplyURI(connectionString) // Connection to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	Collection = client.Database(dbName).Collection(collectionName)
	log.Println("MongoDB Connection instance is ready")
}
