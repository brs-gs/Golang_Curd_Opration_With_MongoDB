package controller

import (
	"encoding/json"
	"golang-Curd-Oprations-With-Mongodb/models"
	"golang-Curd-Oprations-With-Mongodb/utils"
	"log"
	"net/http"
)

var (
	Uc = utils.NewUserController()
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("CreateUser endpoint called")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Failed to decode user: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = Uc.InsertUser(user)
	if err != nil {
		log.Printf("Failed to insert user: %v", err)
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
	log.Println("User created successfully")
}
