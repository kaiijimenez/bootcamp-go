package main

import (
	"fmt"
	"time"
)

func sending(sent chan<- interface{}, value interface{}) {
	fmt.Printf("Value: %v Type of value sending: %T", value, value)
	time.Sleep(time.Second * 1)
	sent <- value
}

func receiving(receive <-chan interface{}) {
	fmt.Println("\nMessage received, preparing to printing...")
	time.Sleep(time.Second * 1)
	fmt.Println("Message received: ", <-receive)
}

func main() {
	value := make(chan interface{})
	go sending(value, 100.12)
	time.Sleep(time.Second * 5)
	receiving(value)
}
