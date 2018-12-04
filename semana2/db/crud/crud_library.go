package crud

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type inMemory interface {
	Create(u *User) (string, error)
	//retrieve(id int) (User, error)
	//update(u *User, id int) (dataBase, error)
	//del(id int) (dataBase, error)
}

type dataBase struct {
	db map[int]*User
}

type User struct {
	Name, Email string
}

type query struct {
	inm inMemory
}

var id = 0

func (db dataBase) Create(u *User) (string, error) {
	id++
	db.db[id] = u
	return fmt.Sprintf("Created: %v", u), nil
}

func (db dataBase) retrieve(id int) (User, error) {
	elem, ok := db.db[id]
	if !ok {
		return User{}, errors.New("Tried to get a key which doesn't exist")
	}
	return *elem, nil
}

func (db dataBase) update(u *User, id int) (dataBase, error) {
	_, ok := db.db[id]
	if !ok {
		return dataBase{}, errors.New("Tried to get a key which doesn't exist")
	}
	db.db[id] = u
	return db, nil
}

func (db dataBase) del(id int) (dataBase, error) {
	_, ok := db.db[id]
	if !ok {
		return dataBase{}, errors.New("Tried to get a key which doesn't exist")
	}
	delete(db.db, id)
	return db, nil
}

func (q *query) Query(s string) {
	str := strings.Fields(s)
	if str[0] == "Create" {
		u := User{str[1], str[2]}
		msg, err := q.inm.Create(&u)
		show(msg, err, "Create")
	}
}

func (db dataBase) showdb(s string) {
	fmt.Printf("DB\n%v\n", s)
	for k, v := range db.db {
		fmt.Printf("Id: %v, User: %v, Email: %v\n", k, v.Name, v.Email)
	}
}

func show(i interface{}, e error, s string) {
	if e != nil {
		log.Fatal(e)
	}
	switch t := i.(type) {
	case User:
		fmt.Println("Retrieved:", i)
	case dataBase:
		t.showdb(s)
	default:
		log.Fatal("Error")
	}
}
