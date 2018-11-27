package main

import "fmt"

func main() {
	any(3)
	any("hi")
	any(1.3)
	any(true)
}

func any(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("Sent: %T\n", v)
	case string:
		fmt.Printf("Sent: %T\n", v)
	case float64:
		fmt.Printf("Sent: %T\n", v)
	default:
		fmt.Printf("Sent: %T\n", v)
	}
}
