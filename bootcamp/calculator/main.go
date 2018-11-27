package main

import (
	"os"

	"github.com/karina.jimenez/bootcamp/calculator/calc"
)

func main() {
	argsWithoutProg := os.Args[1:]
	calc.Calculator(argsWithoutProg)
}
