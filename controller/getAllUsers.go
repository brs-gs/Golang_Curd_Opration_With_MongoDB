package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// This function gets all users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("GetUsers endpoint called")

	allUsers, err := Uc.GetAllUsers()
	if err != nil {
		log.Printf("Failed to get users: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allUsers)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	log.Println("Users retrieved and response sent successfully")
}
