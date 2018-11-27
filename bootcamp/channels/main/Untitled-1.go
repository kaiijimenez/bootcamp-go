func main

import (
	"fmt"
	"time"
)

//func with channel to send

//func with channel to receive and print

//func to stop and return
// the program cannot end until the sender and receiver returned
func main(){
	ch1 := make(chan string)
	
	go func (){
		ch1 <- "ping"
	}()

	msg := <-ch1
	fmt.Prinln(msg)
}