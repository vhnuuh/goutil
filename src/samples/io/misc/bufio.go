// https://segmentfault.com/a/1190000015591319
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var filePath = "./proverbs.txt"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Print(line)
	}
}
