package crud

import (
	"bootcamp-go/semana2/db/errors"
	"fmt"
)

type InMemory interface {
	Create() (string, error)
	Retrieve(id int) (User, error)
	Update(id int) (string, error)
	Del(id int) (string, error)
}

//id para los keys en el map
var id = 0

//User tabla, valores: Name string, Email string
type User struct {
	Name  string
	Email string
}

//db key, values
var db = make(map[int]*User)

func (u *User) Create() (string, error) {
	id++
	_, ok := db[id]
	if !ok {
		db[id] = u
		return fmt.Sprintf("New input\nid: %v, values: %v\nNew: %v\n", id, u, db), nil
	}
	return "", errors.New("duplicated id, try again\n")
}

func Retrieve(id int) (User, error) {
	elem, ok := db[id]
	fmt.Println(&elem, *elem)
	if !ok {
		return User{}, errors.New("id not valid, try again\n")
	}
	return *elem, nil
}

func (to *User) Update(id int) (string, error) {
	elem, ok := db[id]
	if !ok {
		return "", errors.New("id not found, try again\n")
	}
	db[id] = to
	return fmt.Sprintf("Updated!\nOld: %v, New: %v\nNewMap: %v\n", elem, to, db), nil
}

func Del(id int) (string, error) {
	elem, ok := db[id]
	if !ok {
		return "", errors.New("id not found, try again\n")
	}
	delete(db, id)
	return fmt.Sprintf("Deleted: %v, NewMap: %v\n", elem, db), nil
}
