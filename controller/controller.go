package controller

import (
	"github.com/gorilla/mux"
)

func RunController() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user", CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/profile", UpdateExistingUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/{id}", DeleteExistingUser).Methods("DELETE")

	return router

}
