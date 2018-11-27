package main

import (
	"bootcamp-go/semana2/db/crud"
	"fmt"
)

type user struct {
	name, email string
}

var mymap map[int]user

func main() {
	mymap := make(map[int]user)
	mymap[1] = user{"admin", "admin@example.com"}
	u1 := user{
		"admin",
		"admin@example.com",
	}
	created, err := crud.Create(mymap, 2, users.u1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprintf(created)
	r, err := crud.Retrieve(mymap, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprintf(r)
	upd, err := crud.Update(mymap, 2, user{"user3", "user3@example.com"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprintf(upd)
	d, err := crud.Del(mymap, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprintf(d)
}
