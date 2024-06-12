package controller

import (
	"encoding/json"
	"net/http"
)

// This function get all users from Database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allUsers, err := Uc.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(allUsers)
}
