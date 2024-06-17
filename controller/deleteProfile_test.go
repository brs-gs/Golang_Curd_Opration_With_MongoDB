package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDeleteExistingUser(t *testing.T) {
	t.Run("Valid User Deletion", func(t *testing.T) {
		mockCtrl := setupMockUserController(t)
		defer mockCtrl.mockCtrl.Finish()

		userID := "12345"
		mockCtrl.mockUser.EXPECT().DeleteUser(userID).Return(nil)

		req, err := http.NewRequest("DELETE", "/user/{id}", nil)
		assert.NoError(t, err)
		req = mux.SetURLVars(req, map[string]string{"id": userID})
		rr := httptest.NewRecorder()

		DeleteExistingUser(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var response map[string]string
		err = json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "User deleted", response["message"])
		assert.Equal(t, userID, response["userId"])
	})

	t.Run("DeleteUser Failure", func(t *testing.T) {
		mockCtrl := setupMockUserController(t)
		defer mockCtrl.mockCtrl.Finish()

		userID := "12345"
		mockCtrl.mockUser.EXPECT().DeleteUser(userID).Return(errors.New("delete error"))

		req, err := http.NewRequest("DELETE", "/user/{id}", nil)
		assert.NoError(t, err)
		req = mux.SetURLVars(req, map[string]string{"id": userID})
		rr := httptest.NewRecorder()

		DeleteExistingUser(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "delete error")
	})
}
