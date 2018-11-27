package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sending(sent chan<- string, msg string) {
	fmt.Println("Sending message...")
	time.Sleep(time.Second * 1)
	sent <- msg
}

func receiving(receive <-chan string) {
	fmt.Println("Message received, preparing to printing...")
	time.Sleep(time.Second * 1)
	fmt.Println(<-receive)
}

func main() {
	message := make(chan string)
	wg.Add(1)
	go sending(message, "Message to be printed")
	time.Sleep(time.Second * 5)
	receiving(message)
}
