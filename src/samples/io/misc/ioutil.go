package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bytes, err := ioutil.ReadFile("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s", bytes)
}
