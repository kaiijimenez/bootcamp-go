package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x/2
	i := false
	for i == false{
		y := z
		fmt.Println("1.valor de y: ", y)
		z -= (z*z - x) / (2*z)
		fmt.Println("2.valor de z: ", z)
		if y == z{
			i = true
		}	
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
