package main

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber = getSequence()
	fmt.Println(nextNumber())
}
