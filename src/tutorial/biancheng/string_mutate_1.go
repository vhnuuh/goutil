package main

import "fmt"

func mutate(s string) string {
	s[0] = 'a' // any valid unicode character within single quote is a rune
	return s
}

func main() {
	h := "hello, 你好"
	fmt.Println(mutate(h))
}
