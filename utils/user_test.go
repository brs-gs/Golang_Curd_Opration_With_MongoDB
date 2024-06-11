package utils

import (
	"context"
	"golang-Curd-Oprations-With-Mongodb/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserController_InsertUser(t *testing.T) {
	userController := NewUserController()

	user := models.User{
		ID:   primitive.NewObjectID(),
		Name: "test",
		City: "test-city",
		Age:  30,
	}
	userController.InsertUser(user)

	filter := bson.M{"_id": user.ID}
	cursor, err := Collection.Find(context.Background(), filter)
	assert.NoError(t, err)
	defer cursor.Close(context.Background())

	var users []bson.M
	for cursor.Next(context.Background()) {
		var user bson.M
		err := cursor.Decode(&user)
		assert.NoError(t, err)
		users = append(users, user)
	}

	//assert.Len(t, users, 1)
	assert.Equal(t, user.Name, users[1]["name"])
	assert.Equal(t, user.Name, users[1]["city"])
	assert.Equal(t, user.Age, users[1]["age"])
}
