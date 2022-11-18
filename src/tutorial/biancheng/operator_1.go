package main

import "fmt"

func main() {
	var a = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("type of a %T\n", a)
	fmt.Printf("type of b %T\n", b)
	fmt.Printf("type of c %T\n", c)

	ptr = &a
	fmt.Printf("value of a %d\n", a)
	fmt.Printf("value of *ptr %d\n", *ptr)
}
