//https://studygolang.com/articles/19475
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handing is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}
