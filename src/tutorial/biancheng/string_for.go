// for range直接转换为rune
package main

import "fmt"

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func main() {
	name := "你好"
	printCharsAndBytes(name)
}
