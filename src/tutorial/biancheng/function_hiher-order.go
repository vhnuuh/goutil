//https://studygolang.com/articles/12789
package main

import "fmt"

func simple1(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	f := func(a, b int) int {
		return a * b
	}
	simple1(f)

	s := simple2()
	fmt.Println(s(60, 7))
}
