package main

//TODO
//6.ADD DEP INTO THE PROJECT
//7.ADD TESTCASES
//MAKE IT MORE FUNCTIONAL

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kaiijimenez/bootcamp-go/restAPI/handlers"
)

var db *sql.DB

const (
	dbName = "shopping-cart"
	dbPass = "12345"
	dbHost = "localhost"
	dbPort = "33066"
)

func init() {
	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dbSource)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Comiendo tacos.com")
	router := Routes()
	log.Fatal(http.ListenAndServe(":8000", router))

}

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/v1/cart/create", handlers.CreateCart).Methods("POST")
	router.HandleFunc("/v1/cart/add", handlers.AddItems).Methods("POST")
	router.HandleFunc("/v1/cart", handlers.AllItems).Methods("GET")
	router.HandleFunc("/v1/cart/{id}", handlers.UpdateQ).Methods("PATCH")
	router.HandleFunc("/v1/cart/{id}", handlers.DeleteItem).Methods("DELETE")
	router.HandleFunc("/v1/cart", handlers.ClearCart).Methods("DELETE") //clear all cart*/

	return router
}
