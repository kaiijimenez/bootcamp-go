package main

import (
	"fmt"
	"os"
)

//to work with os package, write a file, open a file, read it and close it

// const O_CREATE int = syscall.O_CREAT // create a new file if none exists.

func main() {
	file := "example.csv"
	tw := []byte("hello\nworld\n!")
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		n, e := os.Create(file)
		if e != nil {
			fmt.Println(e)
		}
		_, er := n.Write(tw)
		if e != nil {
			fmt.Println(er)
		}
		r, e := n.Read(tw)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(r)
		n.Close()
	}
	fmt.Println(f)
	f.Close()
}
