package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kaiijimenez/bootcamp-go/restAPI/models"
	"github.com/kaiijimenez/bootcamp-go/restAPI/utils"
)

var (
	items []models.Items
	qt    = "1"
	art   = "http://challenge.getsandbox.com/articles"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	//http.Client with a sensible timeout
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	var c []map[string]string
	var item models.Items

	//Get items from endpoint
	endp, _ := client.Get(art)
	body, _ := ioutil.ReadAll(endp.Body)
	defer endp.Body.Close()
	js := json.Unmarshal(body, &c)
	utils.LogsError("Error trying to unmarshall items: ", js)

	//Inserting items into DB
	for _, v := range c {
		item.ID = v["id"]
		item.Title = v["title"]
		item.Price = v["price"]
		item.Quantity = &qt
		utils.Insert("INSERT INTO shoppingcartdb.items(id, title, price, quantity) VALUES (?,?,?,?);", item, "Error trying to insert into DB")
	}
	i := utils.SelectAll()
	if i == nil {
		log.Fatal("empty db")
		utils.RespondWithError(w, http.StatusNotImplemented, "Empty DB")
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, i)
	return
}

//Response null check SelectAll
func AllItems(w http.ResponseWriter, r *http.Request) {
	i := utils.SelectAll()
	if i == nil {
		utils.RespondWithError(w, http.StatusNotImplemented, "Empty DB")
		log.Print("Empty DB")
	}
	utils.RespondWithJson(w, http.StatusOK, i)
	return
}

func AddItems(w http.ResponseWriter, r *http.Request) {
	//Decoding client request
	var item models.Items
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Fatal(err)
	}

	indb, _ := utils.CheckAvailability(item.ID)
	//If there is no record in the db with same ID then it can be inserted
	if !indb {
		item.Quantity = &qt
		res := utils.Insert("INSERT INTO shoppingcartdb.items (id, title, price, quantity) VALUE (?,?,?,?);", item, "Error trying to insert into DB")
		//As we are adding in to the shopping cart then qt it should be initialize it with 1
		log.Println("Record inserted successfully", res)
		utils.RespondWithJson(w, http.StatusCreated, item)
		return
		//it should end
	}
	//If there is a record with same ID send error response
	utils.RespondWithError(w, http.StatusNotModified, "Record already inserted in DB.")
	return
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]

	indb, item := utils.CheckAvailability(params)
	//If the record is in the database then it can be delete
	if indb {
		del := utils.Delete("DELETE FROM shoppingcartdb.items WHERE id=?", params, "Error trying to delete from DB")
		log.Println("Deleted successfully: ", del)
		utils.RespondWithJson(w, http.StatusNoContent, item) //showing client which item is being deleted
		return
	}
	utils.RespondWithError(w, http.StatusNotFound, "There is not item in DB with same ID")
	return
}

func UpdateQ(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]

	indb, _ := utils.CheckAvailability(params)
	//If there is record in the DB
	if indb {
		//getting quantity request from client
		var item models.Items
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			log.Fatal(err)
		}
		//Updating
		utils.Update("UPDATE shoppingcartdb.items SET quantity = ? WHERE id = ?", params, "Error trying to update the quantity", item.Quantity)
		exist, i := utils.CheckAvailability(params)
		if !exist {
			utils.RespondWithError(w, http.StatusNotFound, "Id doesnt exist in the DB")
			return
		}
		utils.RespondWithJson(w, http.StatusOK, i)
		return
	}
	utils.RespondWithError(w, http.StatusNotFound, "There is not item in DB with same ID")
	return
}

func ClearCart(w http.ResponseWriter, r *http.Request) {
	//verify if items is emtpy to only send a Empty DB
	items := utils.SelectAll()
	if items == nil {
		log.Fatal("empty db")
		utils.RespondWithError(w, http.StatusNotImplemented, "Empty DB")
		return
	}
	if err := utils.CleanDB(); err == nil {
		utils.RespondWithJson(w, http.StatusNoContent, nil)
		return
	}
	utils.RespondWithError(w, http.StatusNotImplemented, "Empty already")
}
