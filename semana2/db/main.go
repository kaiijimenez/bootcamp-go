package main

import (
	"bootcamp-go/semana2/db/crud"
	"fmt"
)

var interfaz crud.InMemory

func main() {
	u1 := crud.User{
		Name:  "admin",
		Email: "admin@example.com",
	}
	u2 := crud.User{
		Name:  "user1",
		Email: "user1@example.com",
	}
	interfaz = u1
	create, err := interfaz.Create()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(create)
	create1, err := u2.Create()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(create1)
	ret, err := crud.Retrieve(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
	upd, err := crud.Update(1, u2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(upd)
	dl, err := crud.Del(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dl)
	u3 := crud.User{
		Name:  "user2",
		Email: "user2@example.com",
	}
	create3, err := u3.Create()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(create3)
}
