package main

import "fmt"

type cb func(int) int

func main() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("annonymous func cb, x = %d\n", x)
		return x
	})
}

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("func cb, x = %d\n", x)
	return x
}
