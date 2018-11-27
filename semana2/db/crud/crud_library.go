package crud

import (
	"bootcamp-go/semana2/db/errors"
	"fmt"
)

type user struct{}

func Create(m map[int]user, id int, u user) (string, error) {
	_, ok := m[id]
	if !ok {
		m[id] = u
		return fmt.Sprintf("New input\nid: %x, values: %x\n New:%x", id, u, m), nil
	}
	return "", errors.New("duplicated id, try again")
}

func Retrieve(m map[int]user, id int) (user, error) {
	elem, ok := m[id]
	if !ok {
		return user{}, errors.New("id not valid, try again")
	}
	return elem, nil
}

func Update(m map[int]user, id int, u user) (string, error) {
	elem, ok := m[id]
	if !ok {
		return "", errors.New("id not found, try again")
	}
	m[id] = u
	return fmt.Sprintf("Updated!\nOld: %x, New: %x", elem, u), nil
}

func Del(m map[int]user, id int) (string, error) {
	elem, ok := m[id]
	if !ok {
		return "", errors.New("id not found, try again")
	}
	delete(m, id)
	return fmt.Sprintf("Deleted: %x, NewMap: %x", elem, m), nil
}
