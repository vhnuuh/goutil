package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " hello, world!"
	fmt.Printf("%s\n", *(&str))
	str1 := strings.TrimSpace(str)
	fmt.Printf("after trimspace: %s\n", str1)
	str2 := strings.Trim(str, "!")
	fmt.Printf("after trim: %s\n", str2)

	str3 := "    ls -l"
	commands := strings.Split(str3, " ")
	fmt.Printf("command: %s\n", strings.TrimSpace(commands[0]))
}
