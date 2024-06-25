package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Function delete the existing user
func DeleteExistingUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("DeleteExistingUser endpoint called")

	params := mux.Vars(r)
	userId := params["id"]
	log.Printf("Attempting to delete user with ID: %s", userId)

	err := Uc.DeleteUser(userId)
	if err != nil {
		log.Printf("Failed to delete user with ID: %s, error: %v", userId, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "User deleted", "userId": userId}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Failed to encode response for user ID: %s, error: %v", userId, err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	log.Printf("User with ID: %s deleted successfully", userId)
}
