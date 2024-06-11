package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Function update the user
func UpdateExistingUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("call in update user function controller************")
	params := mux.Vars(r)
	Uc.UpdateUser(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
