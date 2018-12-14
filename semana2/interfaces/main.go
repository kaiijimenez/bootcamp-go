package main

import (
	"fmt"
	"time"
)

//Example of use interfaces and channels

type stringer interface {
	String() string
}

type myString struct {
	s string
}

func (ms myString) String() string {
	return ms.s
}

func send(sent chan<- string, msg myString) {
	var s stringer
	s = msg
	time.Sleep(time.Second * 1)
	fmt.Println("Message sent")
	sent <- s.String()

}

func receiving(receive <-chan string) {
	fmt.Println("\nMessage received, preparing to printing...")
	time.Sleep(time.Second * 1)
	fmt.Println("Message received: ", <-receive)
}

func main() {
	msg := make(chan string)
	ms := myString{"This is the fly message"}
	go send(msg, ms)
	time.Sleep(time.Second * 5)
	receiving(msg)
	//fmt.Println(msg)
}
