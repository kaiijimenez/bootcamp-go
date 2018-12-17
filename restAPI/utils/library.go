package utils

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kaiijimenez/bootcamp-go/restAPI/models"
)

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@tcp/shoppingcartdb?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return db
}

var items []models.Items

//CheckAvailability returns true if there is an item with same ID in the DB, returns false in case there is not
var CheckAvailability = func(id string) (bool, models.Items) {
	db := dbConn()
	defer db.Close()
	var item models.Items

	row := db.QueryRow("SELECT id FROM shoppingcartdb.items WHERE id=?;", id)
	err := row.Scan(item.ID, item.Title, item.Price, item.Quantity)
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
	var e models.Error
	e.Code = code
	e.Msg = msg
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&items)
}

var LogsError = func(msg string, e error) {
	if e != nil {
		log.Fatal(msg, e)
	}
}

var PrepareExecQuery = func(query string, msgerror string, values interface{}) sql.Result {
	db := dbConn()
	defer db.Close()
	var item models.Items
	q, err := db.Prepare(query)
	LogsError(msgerror, err)
	res, err := q.Exec(item.ID, item.Title, item.Price, item.Quantity)
	LogsError(msgerror, err)
	return res
}

var SelectAll = func() []models.Items {
	db := dbConn()
	defer db.Close()

	var item models.Items

	rows, err := db.Query("SELECT * FROM shoppingcartdb.items;")
	LogsError("Error trying to retrieve the data from db", err)
	if rows == nil {
		return nil
	}
	for rows.Next() {
		rows.Scan(item.ID, item.Title, item.Price, item.Quantity)
		items = append(items, item)
	}
	defer rows.Close()
	return items
}
