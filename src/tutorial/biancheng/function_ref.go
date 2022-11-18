package main

import "fmt"

func main() {
	var a = 100
	var b = 200

	fmt.Printf("before swap, a = %d\n", a)
	fmt.Printf("before swap, b = %d\n", b)

	swap(&a, &b)

	fmt.Printf("after swap, a = %d\n", a)
	fmt.Printf("after swap, b = %d\n", b)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
