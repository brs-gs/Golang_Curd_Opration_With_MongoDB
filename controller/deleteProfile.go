package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Function delete the existing user
func DeleteExistingUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	err := Uc.DeleteUser(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted", "userId": params["id"]})
}
