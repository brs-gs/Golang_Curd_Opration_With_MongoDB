package utils_test

import (
	"golang-Curd-Oprations-With-Mongodb/mocks"
	"golang-Curd-Oprations-With-Mongodb/models"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserController_InsertUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserCtrl := mocks.NewMockUserController(mockCtrl)

	user := models.User{
		Name: "test-name",
		City: "test-city",
		Age:  0,
	}

	mockUserCtrl.EXPECT().InsertUser(user).Return(nil)

	err := mockUserCtrl.InsertUser(user)
	assert.NoError(t, err)
}

func TestUserController_UpdateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserCtrl := mocks.NewMockUserController(mockCtrl)

	userId := "5ecb1c83e2705f56389fdd73"
	updateData := map[string]interface{}{
		"username": "updatedUsername",
	}

	expectedUser := &models.User{
		ID:   primitive.NewObjectID(),
		Name: "updatedUsername",
	}

	mockUserCtrl.EXPECT().UpdateUser(userId, updateData).Return(expectedUser, nil)

	updatedUser, err := mockUserCtrl.UpdateUser(userId, updateData)
	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)
	assert.Equal(t, "updatedUsername", updatedUser.Name)
}

func TestUserController_DeleteUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserCtrl := mocks.NewMockUserController(mockCtrl)

	userId := "5ecb1c83e2705f56389fdd73"

	mockUserCtrl.EXPECT().DeleteUser(userId).Return(nil)

	err := mockUserCtrl.DeleteUser(userId)
	assert.NoError(t, err)
}

func TestUserController_DeleteAllUsers(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserCtrl := mocks.NewMockUserController(mockCtrl)

	mockUserCtrl.EXPECT().DeleteAllUsers(gomock.Any())

	mockUserCtrl.DeleteAllUsers("dummyUserId")
}

func TestUserController_GetAllUsers(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserCtrl := mocks.NewMockUserController(mockCtrl)

	expectedUsers := []primitive.M{
		{
			"_id":   primitive.NewObjectID(),
			"field": "value",
		},
	}

	mockUserCtrl.EXPECT().GetAllUsers().Return(expectedUsers, nil)

	users, err := mockUserCtrl.GetAllUsers()
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Len(t, users, len(expectedUsers))
}
