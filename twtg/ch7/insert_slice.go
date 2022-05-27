package main

import "fmt"

func main() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := InsertStringSlice(s, in, 0)
	fmt.Println(res)
	res = InsertStringSlice(s, in, 3)
	fmt.Println(res)
}

func InsertStringSlice(s, in []string, start int) []string {
	res := make([]string, len(s)+len(in))
	at := copy(res, s[0:start])
	at += copy(res[at:], in)
	copy(res[at:], s[start:])
	return res
}
