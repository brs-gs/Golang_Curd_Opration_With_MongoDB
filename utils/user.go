package utils

import (
	"context"
	"errors"
	"fmt"
	"golang-Curd-Oprations-With-Mongodb/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserController interface {
	InsertUser(user models.User) error
	UpdateUser(userId string, updateData map[string]interface{}) (*models.User, error)
	DeleteUser(userId string) error
	DeleteAllUsers(userId string)
	GetAllUsers() ([]primitive.M, error)
}

type userController struct{}

func NewUserController() UserController {
	return &userController{}
}

func (uc *userController) InsertUser(user models.User) error {
	log.Println("InsertUser called")

	user.ID = primitive.NewObjectID()
	inserted, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Failed to insert user: %v", err)
		return fmt.Errorf("failed to insert user: %w", err)
	}
	log.Printf("Inserted one user in db with id: %v", inserted.InsertedID)
	return nil
}

// update a user record in db
func (uc *userController) UpdateUser(userId string, updateData map[string]interface{}) (*models.User, error) {
	log.Printf("UpdateUser called for userId: %s", userId)

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		return nil, err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateData}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := Collection.FindOneAndUpdate(context.Background(), filter, update, opts)

	if result.Err() != nil {
		log.Printf("Failed to update user: %v", result.Err())
		return nil, result.Err()
	}

	var updatedUser models.User
	if err := result.Decode(&updatedUser); err != nil {
		log.Printf("Failed to decode updated user: %v", err)
		return nil, err
	}
	log.Printf("Modified user: %v", updatedUser)

	return &updatedUser, nil
}

// delete a record from db
func (uc *userController) DeleteUser(userId string) error {
	log.Printf("DeleteUser called for userId: %s", userId)

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		return fmt.Errorf("invalid user ID: %w", err)
	}
	filter := bson.M{"_id": id}

	deleteResult, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return fmt.Errorf("failed to delete user: %w", err)
	}
	if deleteResult.DeletedCount == 0 {
		log.Printf("No user found to delete with ID: %s", userId)
		return errors.New("no user found to delete")
	}
	log.Printf("Deleted user count: %d", deleteResult.DeletedCount)
	return nil
}

// Delete all users from database
func (uc *userController) DeleteAllUsers(userId string) {
	log.Println("DeleteAllUsers called")

	deleteResult, err := Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatalf("Failed to delete all users: %v", err)
	}
	log.Printf("Number of users deleted from db: %d", deleteResult.DeletedCount)
}

// Get all user records from database
func (uc *userController) GetAllUsers() ([]primitive.M, error) {
	log.Println("GetAllUsers called")

	cursor, err := Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Printf("Failed to get users: %v", err)
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	var users []primitive.M
	for cursor.Next(context.Background()) {
		var user bson.M
		if err := cursor.Decode(&user); err != nil {
			log.Printf("Failed to decode user: %v", err)
			return nil, fmt.Errorf("failed to decode user: %w", err)
		}
		users = append(users, user)
	}

	if err := cursor.Close(context.Background()); err != nil {
		log.Printf("Failed to close cursor: %v", err)
		return nil, fmt.Errorf("failed to close cursor: %w", err)
	}

	log.Println("Users retrieved successfully")
	return users, nil
}
