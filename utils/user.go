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

	user.ID = primitive.NewObjectID()
	inserted, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	fmt.Println("Inserted one user in db with id: ", inserted.InsertedID)
	return nil
}

// update a user record in db
func (uc *userController) UpdateUser(userId string, updateData map[string]interface{}) (*models.User, error) {

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateData}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := Collection.FindOneAndUpdate(context.Background(), filter, update, opts)

	if result.Err() != nil {
		return nil, result.Err()
	}

	var updatedUser models.User
	if err := result.Decode(&updatedUser); err != nil {
		return nil, err
	}
	fmt.Println("Modified user: ", updatedUser)

	return &updatedUser, nil
}

// delete a record from db
func (uc *userController) DeleteUser(userId string) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}
	filter := bson.M{"_id": id}

	deleteCount, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	if deleteCount.DeletedCount == 0 {
		return errors.New("no user found to delete")
	}
	fmt.Println("Delete count is:", deleteCount.DeletedCount)
	return nil
}

// Delete all user from database
func (uc *userController) DeleteAllUsers(userId string) {

	deleteResult, err := Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of user are deleted from db is: ", deleteResult.DeletedCount)
}

// Get all user records from database
func (uc *userController) GetAllUsers() ([]primitive.M, error) {
	cursor, err := Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	var users []primitive.M
	for cursor.Next(context.Background()) {
		var user bson.M
		err := cursor.Decode(&user)
		if err != nil {
			return nil, fmt.Errorf("failed to decode user: %w", err)
		}
		users = append(users, user)
	}
	if err := cursor.Close(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to close cursor: %w", err)
	}
	return users, nil
}
