package controller

import (
	"encoding/json"
	"errors"
	"golang-Curd-Oprations-With-Mongodb/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUsers(t *testing.T) {
	t.Run("Get All Users Successfully", func(t *testing.T) {
		mockCtrl := setupMockUserController(t)
		defer mockCtrl.mockCtrl.Finish()

		// Adjust mock return to match expected type
		users := []primitive.M{
			{"name": "test-user12"},
			{"name": "test-user13"},
		}
		mockCtrl.mockUser.EXPECT().GetAllUsers().Return(users, nil)

		req, err := http.NewRequest("GET", "/users", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		GetUsers(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var responseUsers []models.User
		err = json.NewDecoder(rr.Body).Decode(&responseUsers)
		assert.NoError(t, err)

		// Convert mock return type to expected response type
		expectedUsers := []models.User{
			{Name: "test-user12"},
			{Name: "test-user13"},
		}
		assert.Equal(t, expectedUsers, responseUsers)
	})

	t.Run("Get All Users Failure", func(t *testing.T) {
		mockCtrl := setupMockUserController(t)
		defer mockCtrl.mockCtrl.Finish()

		mockCtrl.mockUser.EXPECT().GetAllUsers().Return(nil, errors.New("database error"))

		req, err := http.NewRequest("GET", "/users", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		GetUsers(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "database error")
	})
}
