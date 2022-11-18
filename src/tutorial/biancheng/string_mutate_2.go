package main

import (
	"fmt"
	"unicode/utf8"
)

func mutate(s []rune) string {
	s[0] = 'a' // any valid unicode character within single quote is a rune
	return string(s)
}

func main() {
	h := "hello, 你好"
	//fmt.Println(mutate(h))
	fmt.Println(mutate([]rune(h)))
	fmt.Println("length of %s is %d\n", h, utf8.RuneCountInString(h))
}
