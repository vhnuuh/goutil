package main

import "fmt"

func main() {
	var sum = 17
	var count = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean's value is: %f\n", mean)
}
