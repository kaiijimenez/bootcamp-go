package main

import (
	"fmt"
)

const diff = 1e-6

type ErrNegativaSqrt float64

func (e ErrNegativaSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativaSqrt(x)
	}
	z := x
	oldz := 0.0
	for {
		if v := z - oldz; -diff < v && v < diff {
			return z, nil
		} else {
			oldz = z
			z -= (z*z - x) / (2 * z)
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
