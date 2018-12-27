package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kaiijimenez/bootcamp-go/restAPI/models"
)

func DbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@tcp(shoppingdb)/shoppingcartdb?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return db
}

var items []models.Items

//CheckAvailability returns true if there is an item with same ID in the DB, returns false in case there is not
var CheckAvailability = func(id string) (bool, models.Items) {
	db := DbConn()
	defer db.Close()
	var item models.Items

	row := db.QueryRow("SELECT id FROM shoppingcartdb.items WHERE id=?;", id)
	err := row.Scan(&item.ID, &item.Title, &item.Price, &item.Quantity)
	fmt.Println(item)
	if err == sql.ErrNoRows {
		//in case there is not an item with the same ID
		return false, item
	}
	//in case there is an item with the same ID
	return true, item
}

var RespondWithJson = func(w http.ResponseWriter, code int, items interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&items)
}

var RespondWithError = func(w http.ResponseWriter, code int, msg string) {
	e := make(map[string]interface{})
	e["code"] = code
	e["message"] = msg
	fmt.Println(e)
	RespondWithJson(w, code, e)
}

var LogsError = func(msg string, e error) {
	if e != nil {
		log.Fatal(msg, e)
	}
}

var Insert = func(query string, values models.Items, msgerror string) sql.Result {
	db := DbConn()
	defer db.Close()
	q, err := db.Prepare(query)
	LogsError(msgerror, err)
	res, err := q.Exec(values.ID, values.Title, values.Price, values.Quantity)
	LogsError(msgerror, err)
	return res
}

var Delete = func(query, id, msgerror string) sql.Result {
	db := DbConn()
	defer db.Close()
	q, err := db.Prepare(query)
	LogsError(msgerror, err)
	res, err := q.Exec(id)
	LogsError(msgerror, err)
	return res
}

var Update = func(query, id, msgerror string, qt *string) sql.Result {
	db := DbConn()
	defer db.Close()
	q, err := db.Prepare(query)
	LogsError(msgerror, err)
	res, err := q.Exec(qt, id)
	LogsError(msgerror, err)
	return res
}

var SelectAll = func() []models.Items {
	db := DbConn()
	defer db.Close()
	rows, err := db.Query("SELECT *  FROM shoppingcartdb.items")
	LogsError("Error trying to retrieve the data from db", err)
	if rows == nil {
		return nil
	}
	for rows.Next() {
		var item models.Items
		rows.Scan(&item.ID, &item.Title, &item.Price, &item.Quantity)
		items = append(items, item)
	}
	return items
}

func CleanDB() error {
	db := DbConn()
	defer db.Close()
	_, err := db.Query("DELETE FROM shoppingcartdb.items;")
	if err != nil {
		LogsError("Not able to delete from items table", err)
		return err
	}
	return nil
}
