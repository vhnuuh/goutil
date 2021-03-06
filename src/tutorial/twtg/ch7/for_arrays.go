package main

import "fmt"

func main() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}
	for i, x := range arr1 {
		fmt.Println("arr1 item", i, "is", arr1[i], x)
	}

	a := [...]string{"a", "b", "c", "d"}
	fmt.Printf("arr1 type %T, a type %T", arr1, a)
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}
