package main

import (
	"bootcamp-go/semana3/os/examples"
)

//create(int, string) map[int]string
var (
	w = "write.json"
	r = "read.json"
)

func main() {
	//write a json file
	mp := examples.JsonStruct{
		JStruct: make(map[int]string),
	}
	mp.Add(1, "one", w)
	mp.Add(2, "two", w)
	mp.Add(3, "three", w)
	mp.Add(4, "four", w)
	mp.Add(5, "five", w)
	//mp.Add(1, "other-one", w)

	//reading a json file
	examples.GetData(r)
}
