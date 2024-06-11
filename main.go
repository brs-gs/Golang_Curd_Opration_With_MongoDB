package main

import (
	"fmt"
	"golang-Curd-Oprations-With-Mongodb/controller"
	"log"
	"net/http"
)

func main() {
	r := controller.RunController()
	fmt.Println("Server is getting started...")
	fmt.Println("Listening at port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))

}
