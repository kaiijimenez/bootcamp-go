package main

import "fmt"

//goroutine to get the squares
func squares(num int, sqr chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	sqr <- sum
}

//goroutine to get cubes
func cubes(num int, cube chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	cube <- sum
}

//main function to get the sum of the squares + cubes
func main() {
	number := 589
	sqr := make(chan int)
	cub := make(chan int)
	go squares(number, sqr)
	go cubes(number, cub)
	sqrRes, cubRes := <-sqr, <-cub
	fmt.Println("Result: ", sqrRes+cubRes)
}
