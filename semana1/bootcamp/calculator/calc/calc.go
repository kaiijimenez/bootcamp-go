package calc

import (
	"errors"
	"fmt"
	"strconv"
)

func Calculator(input []string) {
	num1, _ := strconv.ParseFloat(input[0], 64)
	oper := input[1]
	num2, _ := strconv.ParseFloat(input[2], 64)

	switch oper {
	case "+":
		fmt.Println("Sum :", num1+num2)
	case "-":
		fmt.Println("Substraction :", num1-num2)
	case "*":
		fmt.Println("Multiplication :", num1*num2)
	default:
		num, err := zeroDivisionError(num1, num2)
		if err != nil {
			fmt.Println("Failed: ", err)
		} else {
			fmt.Println("Division:", num)
		}
	}
}

func zeroDivisionError(y, x float64) (float64, error) {
	if x == 0 {
		return 0, errors.New("ZeroDivisionError, Cannot divide under 0")
	}
	return y / x, nil
}
