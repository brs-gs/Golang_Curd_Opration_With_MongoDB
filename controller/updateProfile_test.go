package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang-Curd-Oprations-With-Mongodb/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestUpdateExistingUser(t *testing.T) {
	t.Run("Valid Update User", func(t *testing.T) {
		mockCtrl := setupMockUserController(t)
		defer mockCtrl.mockCtrl.Finish()

		userID := "507f1f77bcf86cd799439011"
		updateData := map[string]interface{}{
			"name": "Updated Name",
			"city": "Updated City",
		}
		updatedUser := models.User{
			Name: "Updated Name",
			City: "Updated City",
		}
		updateDataJson, _ := json.Marshal(updateData)

		mockCtrl.mockUser.EXPECT().UpdateUser(userID, updateData).Return(&updatedUser, nil)

		req, err := http.NewRequest("PUT", "/user/{id}", bytes.NewBuffer(updateDataJson))
		assert.NoError(t, err)
		req = mux.SetURLVars(req, map[string]string{"id": userID})
		rr := httptest.NewRecorder()

		UpdateExistingUser(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var responseUser models.User
		err = json.NewDecoder(rr.Body).Decode(&responseUser)
		assert.NoError(t, err)
		assert.Equal(t, updatedUser, responseUser)
	})

	t.Run("Invalid Request Payload", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/user/{id}", bytes.NewBuffer([]byte("invalid json")))
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		UpdateExistingUser(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "Invalid request payload")
	})

	t.Run("UpdateUser Failure", func(t *testing.T) {
		mockCtrl := setupMockUserController(t)
		defer mockCtrl.mockCtrl.Finish()

		userID := "507f1f77bcf86cd799439011"
		updateData := map[string]interface{}{
			"name": "Updated Name",
			"city": "Updated City",
		}
		updateDataJson, _ := json.Marshal(updateData)

		mockCtrl.mockUser.EXPECT().UpdateUser(userID, updateData).Return(nil, errors.New("update error"))

		req, err := http.NewRequest("PUT", "/user/{id}", bytes.NewBuffer(updateDataJson))
		assert.NoError(t, err)
		req = mux.SetURLVars(req, map[string]string{"id": userID})
		rr := httptest.NewRecorder()

		UpdateExistingUser(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "update error")
	})
}
