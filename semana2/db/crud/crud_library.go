package crud

import (
	"bootcamp-go/semana2/db/errors"
	"fmt"
)

type InMemory interface {
	Create() (string, error)
}

//id para los keys en el map
var id = 0

//User tabla, valores: Name string, Email string
type User struct {
	Name  string
	Email string
}

//db key, values
var db = make(map[int]User)

func (u User) Create() (string, error) {
	id++
	_, ok := db[id]
	if !ok {
		db[id] = u
		return fmt.Sprintf("New input\nid: %v, values: %v\nNew: %v", id, u, db), nil
	}
	return "", errors.New("duplicated id, try again")
}

func Retrieve(id int) (User, error) {
	elem, ok := db[id]
	if !ok {
		return User{}, errors.New("id not valid, try again")
	}
	return elem, nil
}

func Update(id int, to User) (string, error) {
	elem, ok := db[id]
	if !ok {
		return "", errors.New("id not found, try again")
	}
	db[id] = to
	return fmt.Sprintf("Updated!\nOld: %v, New: %v\nNewMap: %v", elem, to, db), nil
}

func Del(id int) (string, error) {
	elem, ok := db[id]
	if !ok {
		return "", errors.New("id not found, try again")
	}
	delete(db, id)
	return fmt.Sprintf("Deleted: %v, NewMap: %v", elem, db), nil
}
