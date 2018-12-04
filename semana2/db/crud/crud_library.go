package crud

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type InMemory interface {
	Create(u *User) (string, error)
	Retrieve(id int) (string, error)
	Update(u *User, id int) (string, error)
	Del(id int) (string, error)
}

type Database struct {
	DB map[int]*User
}

type User struct {
	Name, Email string
}

type Query struct {
	Inm InMemory
}

var id = 0

func (db Database) Create(u *User) (string, error) {
	id++
	db.DB[id] = u
	return fmt.Sprintf("Created: %v\nIn DB: %v", u, db.Showdb()), nil
}

func (db Database) Retrieve(id int) (string, error) {
	elem, ok := db.DB[id]
	if !ok {
		return "", errors.New("Tried to get a key which doesn't exist")
	}
	return fmt.Sprintf("Value from map: %v", elem), nil
}

func (db Database) Update(u *User, id int) (string, error) {
	elem, ok := db.DB[id]
	if !ok {
		return "", errors.New("Tried to get a key which doesn't exist")
	}
	db.DB[id] = u
	return fmt.Sprintf("Updated, old:%v, new:%v\nNew map: %v", elem, u, db.Showdb()), nil
}

func (db Database) Del(id int) (string, error) {
	elem, ok := db.DB[id]
	if !ok {
		return "", errors.New("Tried to get a key which doesn't exist")
	}
	delete(db.DB, id)
	return fmt.Sprintf("Deleted: %v\nIn DB: %v", elem, db.Showdb()), nil
}

// GetQuery get the query from user and apply the proper functionality
func (q *Query) GetQuery(s string) (string, error) {
	str := strings.Fields(s)
	var (
		msg string
		err error
	)
	switch str[0] {
	case "Create":
		u := User{str[1], str[2]}
		msg, err = q.Inm.Create(&u)
	case "Update":
		u := User{str[3], str[4]}
		id, _ := strconv.Atoi(str[2])
		msg, err = q.Inm.Update(&u, id)
	case "Get":
		id, _ := strconv.Atoi(str[2])
		msg, err = q.Inm.Retrieve(id)
	case "Delete":
		id, _ := strconv.Atoi(str[2])
		msg, err = q.Inm.Del(id)
	default:
		msg = "Error"
		err = nil
	}
	return msg, err
}

func (db Database) Showdb() map[int]User {
	data := make(map[int]User)
	if db.DB == nil {
		return data
	}
	for k, v := range db.DB {
		data[k] = *v
	}
	return data
}

func show(s string, e error) {
	if e != nil {
		log.Fatal(e)
	} else {
		fmt.Println(s)
	}
}

func Run() {
	q := Query{
		Inm: Database{
			DB: make(map[int]*User),
		},
	}
	cre, err := q.GetQuery("Create admin admin@example.com")
	show(cre, err)
	cre1, err := q.GetQuery("Create user user@example.com")
	show(cre1, err)
	upd, err := q.GetQuery("Update id 1 admin1 admin1@example.com")
	show(upd, err)
	get, err := q.GetQuery("Get id 1")
	show(get, err)
	de, err := q.GetQuery("Delete id 1")
	show(de, err)
	//RUN THE QUERIES FOR THE INMEMORY OF THE DB
}
