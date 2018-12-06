package main

import (
	"bootcamp-go/semana3/os/examples"
	"log"
)

//create(int, string) map[int]string
var (
	f = "file.txt"
	b []byte
)

func main() {
	mp := examples.Data{
		M: make(map[int]string),
	}
	mp.Add(3, "three", f)
	e := examples.Load(f)
	if e != nil {
		log.Fatal(e)
	}
}
