package controller

import (
	"encoding/json"
	"net/http"
)

// This function get all users from Database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allUser := Uc.GetAllUsers()
	json.NewEncoder(w).Encode(allUser)

}
