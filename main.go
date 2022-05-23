package main

import (
	"fmt"
	"gitlab/moxdata/forecast-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	log.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/forecasts", controllers.GetForecastsFor).Methods("GET")
	myRouter.Queries("start-date",
		"{start-date:20\\d\\d-[01]?\\d-[0123]\\d}",
		"end-date", "{end-date:20\\d\\d-[01]?\\d-[0123]\\d",
		"city", "{city:^([a-z]+|[A-Z]+){1}")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
