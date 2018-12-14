package main

import (
	"os"

	"github.com/kaiijimenez/bootcamp-go/semana1/bootcamp/calculator/calc"
)

func main() {
	argsWithoutProg := os.Args[1:]
	calc.Calculator(argsWithoutProg)
}
