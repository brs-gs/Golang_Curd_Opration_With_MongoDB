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
	fmt.Println("Listening at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))

}
