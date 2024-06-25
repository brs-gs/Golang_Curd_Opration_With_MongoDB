package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Function update the user
func UpdateExistingUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("UpdateExistingUser endpoint called")

	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		log.Printf("Failed to decode update data: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	userId := params["id"]
	log.Printf("Attempting to update user with ID: %s", userId)

	updatedUser, err := Uc.UpdateUser(userId, updateData)
	if err != nil {
		log.Printf("Failed to update user with ID: %s, error: %v", userId, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		log.Printf("Failed to encode response for user ID: %s, error: %v", userId, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("User with ID: %s updated successfully", userId)
}
