package main


import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kaiijimenez/bootcamp-go/restAPI/handlers"
)

func main() {
	fmt.Println("Running")
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
