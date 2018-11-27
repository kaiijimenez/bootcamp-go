package main

import (
	"fmt"
)

type user struct {
	name, email string
}

var mymap map[int]user

func main() {
	mymap := make(map[int]user)
	mymap[1] = user{
		"a", "b",
	}
	u1 := user{
		"user1",
		"user1@example.com",
	}
	create(mymap, 2, u1)
	retrieve(mymap, 2)
	update(mymap, 1, u1)
	del(mymap, 1)

}

func create(m map[int]user, id int, u user) {
	m[id] = u
	fmt.Println(m)
}

func retrieve(m map[int]user, id int) {
	valores := m[id]
	fmt.Println(valores)
}

func update(m map[int]user, id int, u user) {
	m[id] = u
	fmt.Println(m)
}

func del(m map[int]user, id int) {
	delete(m, id)
	fmt.Println(m)
}
