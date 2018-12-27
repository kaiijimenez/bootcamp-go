package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaiijimenez/bootcamp-go/earthquakes/controllers"
)

func main() {
	fmt.Println("RUNNING...")
	router := mux.NewRouter()
	//get an eartquake by id
	router.HandleFunc("/v1/earthquake/{key}", controllers.GetEarthquake).Methods("GET")
	//delete an eartquake by id
	router.HandleFunc("/v1/earthquake/{key}", controllers.DeleteFromMemory).Methods("DELETE")
	//update an eartquake by id
	router.HandleFunc("/v1/earthquake/{key}", controllers.UpdateEarthquake).Methods("PUT") // OR PATCH
	//get an earthquake from a period of two days past
	router.HandleFunc("/v1/earthquake/2/days", controllers.GetFromPeriod).Methods("GET")
	//report a new earthquake
	router.HandleFunc("/v1/earthquake/report", controllers.ReportEarthquake).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
