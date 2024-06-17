package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang-Curd-Oprations-With-Mongodb/mocks"
	"golang-Curd-Oprations-With-Mongodb/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// Mocking UserController
type MockUserController struct {
	mockCtrl *gomock.Controller
	mockUser *mocks.MockUserController
}

func setupMockUserController(t *testing.T) *MockUserController {
	ctrl := gomock.NewController(t)
	mockUser := mocks.NewMockUserController(ctrl)
	Uc = mockUser
	return &MockUserController{
		mockCtrl: ctrl,
		mockUser: mockUser,
	}
}

func TestCreateUser(t *testing.T) {
	t.Run("Valid User Creation", func(t *testing.T) {
		mockCtrl := setupMockUserController(t)
		defer mockCtrl.mockCtrl.Finish()

		user := models.User{
			Name: "test-user11",
			City: "test-city11",
			Age:  0,
		}
		userJson, _ := json.Marshal(user)

		mockCtrl.mockUser.EXPECT().InsertUser(user).Return(nil)

		req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(userJson))
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		CreateUser(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		var responseUser models.User
		err = json.NewDecoder(rr.Body).Decode(&responseUser)
		assert.NoError(t, err)
		assert.Equal(t, user, responseUser)
	})

	t.Run("Invalid Request Payload", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/user", bytes.NewBuffer([]byte("invalid json")))
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		CreateUser(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "Invalid request payload")
	})

	t.Run("InsertUser Failure", func(t *testing.T) {
		mockCtrl := setupMockUserController(t)
		defer mockCtrl.mockCtrl.Finish()

		user := models.User{
			Name: "test-user11",
			City: "test-city",
			Age:  0}
		userJson, _ := json.Marshal(user)

		mockCtrl.mockUser.EXPECT().InsertUser(user).Return(errors.New("insert error"))

		req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(userJson))
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		CreateUser(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "Failed to insert user")
	})
}
