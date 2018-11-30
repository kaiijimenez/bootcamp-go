package main

import (
	"bootcamp-go/semana2/db/crud"
	"fmt"
)

type users []crud.User

var imp crud.InMemory

func main() {
	u := crud.User{Name: "user2", Email: "user2@example.com"}
	users := []crud.User{
		crud.User{Name: "admin", Email: "admin@example.com"},
		crud.User{Name: "user1", Email: "user1@example.com"},
	}
	for i := range users {
		//imp := &users[i]
		create, err := users[i].Create()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(create)
	}
	fmt.Println("%v, %T", imp, imp)
	ret, err := crud.Retrieve(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ret: ", ret)
	upd, err := u.Update(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(upd)
	dl, err := imp.Del(2)
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
