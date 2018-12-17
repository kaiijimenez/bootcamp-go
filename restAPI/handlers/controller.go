package handlers

import (
	"encoding/json"
	"fmt"
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

	var c []models.Cart
	var d []map[string]string

	//Get items from endpoint
	endp, _ := client.Get(art)
	body, _ := ioutil.ReadAll(endp.Body)
	defer endp.Body.Close()
	js := json.Unmarshal(body, &c)
	e := json.Unmarshal(body, &d)
	utils.LogsError("Error trying to unmarshall items: ", js)
	utils.LogsError("Error trying to unmarshall items: ", e)
	fmt.Println(c)
	fmt.Println(d)

	//Inserting items into DB
	for _, v := range c {
		v.Prod.Quantity = &qt
		res := utils.PrepareExecQuery("INSERT INTO shoppingcartdb.items(id, title, price, quantity) VALUES (?,?,?,?);", "Error trying to insert into DB", v)
		log.Println("Successfully inserted: ", res)
	}
	i := utils.SelectAll()
	if i == nil {
		log.Fatal("empty db")
		utils.RespondWithError(w, http.StatusNotImplemented, "Empty DB")
	}
	utils.RespondWithJson(w, http.StatusCreated, i)
}

//Verify if the table is empty so it should send a diff response
func AllItems(w http.ResponseWriter, r *http.Request) {
	i := utils.SelectAll()
	if i == nil {
		log.Fatal("empty db")
		utils.RespondWithError(w, http.StatusNotImplemented, "Empty DB")
	}
	utils.RespondWithJson(w, http.StatusOK, i)
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
		insert := utils.PrepareExecQuery("INSERT INTO shoppingcartdb.items (id, title, price, quantity) VALUE (?,?,?,?);", "Error trying to insert into DB")
		//As we are adding in to the shopping cart then qt it should be initialize it with 1
		res, err := insert.Exec(item.ID, item.Title, item.Price, &qt)
		utils.LogsError("Error executing values in the insert query: ", err)
		log.Println("Record inserted successfully", res)
		utils.RespondWithJson(w, http.StatusCreated, item)
		//it should end
	}
	//If there is a record with same ID send error response
	utils.RespondWithError(w, http.StatusNotModified, "Record already inserted in DB.")
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]

	indb, item := utils.CheckAvailability(params)
	//If the record is in the database then it can be delete
	if indb {
		del := utils.PrepareExecQuery("DELETE FROM shoppingcartdb.item WHERE id=?", "Error trying to delete from DB")
		res, err := del.Exec(params)
		utils.LogsError("Error trying to execute the delete query: ", err)
		log.Println("Deleted successfully: ", res)
		utils.RespondWithJson(w, http.StatusNoContent, item) //showing client which item is being deleted
	}
	utils.RespondWithError(w, http.StatusNotFound, "There is not item in DB with same ID")
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
		upd := utils.PrepareExecQuery("UPDATE shoppingcartdb.item SET quantity = ? WHERE id = ?", "Error trying to update the quantity")
		res, err := upd.Exec(item.Quantity, params)
		utils.LogsError("Error trying to execute the update query: ", err)
		log.Println("UPDATED successfully", res)
		_, i := utils.CheckAvailability(params) //supposed that it should return me the record updated
		utils.RespondWithJson(w, http.StatusOK, i)
		return
	}
	utils.RespondWithError(w, http.StatusNotFound, "There is not item in DB with same ID")
}

func ClearCart(w http.ResponseWriter, r *http.Request) {
	//verify if items is emtpy to only send a Empty DB
	items := utils.SelectAll()
	if items == nil {
		log.Fatal("empty db")
		utils.RespondWithError(w, http.StatusNotImplemented, "Empty DB")
	}
	for _, values := range items {
		del := utils.PrepareExecQuery("DELETE FROM shoppingcartdb.item WHERE id=?", "Error trying to delete from DB")
		res, err := del.Exec(values.ID)
		utils.LogsError("Error trying to execute delete query (Clear method): ", err)
		log.Println("Deleted successfully: ", res)
		utils.RespondWithJson(w, http.StatusNoContent, values)
	}

}
