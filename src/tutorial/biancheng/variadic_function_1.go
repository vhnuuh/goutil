package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found int ", nums)
	}
	fmt.Printf("\n")
}

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 100)
	find(87)
}
