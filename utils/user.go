package utils

import (
	"context"
	"fmt"
	"golang-Curd-Oprations-With-Mongodb/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController interface {
	InsertUser(user models.User)
	UpdateUser(userId string)
	DeleteUser(userId string)
	DeleteAllUsers(userId string)
	GetAllUsers() []primitive.M
}

type userController struct{}

func NewUserController() UserController {
	return &userController{}
}

// insert a user in db
func (uc *userController) InsertUser(user models.User) {
	user.ID = primitive.NewObjectID()
	inserted, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one user in db with id: ", inserted.InsertedID)
}

// update a user record in db
func (uc *userController) UpdateUser(userId string) {

	fmt.Println("call in update user function")

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"city": "Nashik"}}

	result, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mdified count is: ", result.ModifiedCount)

}

// delete a record from db
func (uc *userController) DeleteUser(userId string) {

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}

	deleteCount, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete count is: ", deleteCount.DeletedCount)
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
func (uc *userController) GetAllUsers() []primitive.M {

	cursor, err := Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var users []primitive.M
	for cursor.Next(context.Background()) {
		var user bson.M
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
		defer cursor.Close(context.Background())

	}
	return users

}
