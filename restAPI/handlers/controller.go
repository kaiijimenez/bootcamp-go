package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/kaiijimenez/bootcamp-go/restAPI/models"

	"github.com/gorilla/mux"
)

var (
	items []models.Items
	qt    = "1"
	art   = "http://challenge.getsandbox.com/articles"
)

//CreateCart with the values from third party uri
func CreateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	var c []map[string]string
	var cart models.Items
	endp, _ := client.Get(art)
	body, _ := ioutil.ReadAll(endp.Body)
	defer endp.Body.Close()
	if js := json.Unmarshal(body, &c); js != nil {
		log.Fatal(js)
	}
	for _, v := range c {
		cart.ID = v["id"]
		cart.Title = v["title"]
		cart.Price = v["price"]
		cart.Quantity = &qt
		items = append(items, cart)
	}

	w.WriteHeader(http.StatusCreated) //Created
	fmt.Println(items)
	json.NewEncoder(w).Encode(&items)
}

func AllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //Good Response
	json.NewEncoder(w).Encode(&items)
}

//CANNOT ADD MORE THAN ONE ITEM
func AddItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Items
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Fatal(err)
	}
	for _, v := range items {
		//In case the ID is already created
		if item.ID == v.ID {
			w.WriteHeader(http.StatusNotModified) //client has the response already
			return
		}
	}
	//As we are adding in to the shopping cart then qt it should be initialize it with 1
	w.WriteHeader(http.StatusCreated)
	item.Quantity = &qt
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]
	ver, index := CheckAvailability(params)
	if ver {
		items = append(items[:index], items[index+1:]...)
		w.WriteHeader(http.StatusNoContent) //when deleting
		json.NewEncoder(w).Encode(items)
		return
	}
	w.WriteHeader(http.StatusNotFound) //in case does not exist
	return

}

//CHECK AS FOR NOW WORKING AS EXPECTED BUT I HAVE TO DELETE THE ITEM AND THEN ADD IT AGAIN IS THIS FUNCTIONAL?
func UpdateQ(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]
	var item models.Items
	for key, value := range items {
		if value.ID == params {
			items = append(items[:key], items[key+1:]...)
			if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
				log.Fatal(err)
			}
			// it should update only the quantity of an EXISTING
			value.Quantity = item.Quantity
			items = append(items, value)
			json.NewEncoder(w).Encode(items)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	return
}

func ClearCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if items != nil {
		items = nil
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(items)
		return
	}
	w.WriteHeader(http.StatusNotFound)

}

func CheckAvailability(id string) (bool, int) {
	for k, v := range items {
		if v.ID == id {
			return true, k
		}
	}
	return false, 0
}
