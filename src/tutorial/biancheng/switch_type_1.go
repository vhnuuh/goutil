package main

import "fmt"

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x: %T", i)
	case int:
		fmt.Printf("type of x is int")
	case float64:
		fmt.Printf("type of x is float64")
	case func(int) float64:
		fmt.Printf("type of x is func(int) float64")
	default:
		fmt.Printf("unknown")
	}
}
