package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Function update the user
func UpdateExistingUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	updatedUser, err := Uc.UpdateUser(params["id"], updateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
