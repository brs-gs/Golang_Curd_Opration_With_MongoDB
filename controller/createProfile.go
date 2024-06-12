package controller

import (
	"encoding/json"
	"golang-Curd-Oprations-With-Mongodb/models"
	"golang-Curd-Oprations-With-Mongodb/utils"
	"net/http"
)

var (
	Uc = utils.NewUserController()
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = Uc.InsertUser(user)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
