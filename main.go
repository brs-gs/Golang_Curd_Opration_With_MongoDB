package main

import (
	"fmt"
	"golang-Curd-Oprations-With-Mongodb/controller"

	httpSwagger "github.com/swaggo/http-swagger"

	"log"
	"net/http"
)

func main() {
	r := controller.RunController()
	fmt.Println("Server is getting started...")
	fmt.Println("Listening at port 8000...")
	//swagger endpoint
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8000", r))

}
