package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Function delete the existing user
func DeleteExistingUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parms := mux.Vars(r)
	Uc.DeleteUser(parms["id"])
	json.NewEncoder(w).Encode(parms["id"])

}
