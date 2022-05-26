package main

import "fmt"

var s []int

func main() {
	s = []int{1, 2, 3}
	fmt.Println("The length of s before enlarging is:", len(s))
	fmt.Println(s)
	s = enlarge(s, 5)
	fmt.Println("The length of s after enlarging is:", len(s))
	fmt.Println(s)
}

func enlarge(src []int, factor int) []int {
	res := make([]int, len(src)*factor)
	copy(res, src)
	return res
}
