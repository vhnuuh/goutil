package main

import "fmt"

func main() {
	var a = 1
	var ptr1 = &a
	var ptr2 = &ptr1
	var ptr3 = &ptr2

	fmt.Println("a: ", a)
	fmt.Println("ptr1: ", ptr1)
	fmt.Println("ptr2: ", ptr2)
	fmt.Println("ptr3: ", ptr3)
	fmt.Println("*ptr1: ", *ptr1)
	fmt.Println("**ptr2: ", **ptr2)
	fmt.Println("**(*ptr3): ", **(*ptr3))
}
