package main

import (
	_ "golang-Curd-Oprations-With-Mongodb/docs"
	"log"
	"net/http"

	"golang-Curd-Oprations-With-Mongodb/controller"
)

// @title golang-Curd-Oprations-With-Mongodb
// @version 1.0
// @description This is a sample server for golang-Curd-Oprations-With-Mongodb.
// @host localhost:8000
// @BasePath /
func main() {
	r := controller.RunController()
	log.Println("Server is getting started...")
	log.Println("Listening at port 8000...")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
