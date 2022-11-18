package main

import "fmt"

func Factorial(n uint64) uint64 {
	if n > 0 {
		result := n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	var i = 15
	fmt.Printf("%d result is %d\n", i, Factorial(uint64(i)))
}
