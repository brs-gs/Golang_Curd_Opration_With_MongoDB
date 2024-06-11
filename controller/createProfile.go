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

// This function Create New user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	Uc.InsertUser(user)
	json.NewEncoder(w).Encode(user)

}
