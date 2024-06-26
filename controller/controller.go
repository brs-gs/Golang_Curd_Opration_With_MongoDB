package controller

import (
	"log"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RunController() *mux.Router {
	log.Println("Starting RunController")
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user", CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", UpdateExistingUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/{id}", DeleteExistingUser).Methods("DELETE")
	// swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("RunController initialized")
	return router
}
