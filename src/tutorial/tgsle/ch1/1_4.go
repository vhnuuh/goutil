package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

func sampleReadSlice() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
}

func errReadSlice() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com"), 16)
	line, err := reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
	line, err = reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
}

func sampleReadBytes() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home\n of gophers"))
	line, _ := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
}

func samplePeek() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
	go Peek(reader)
	go reader.ReadBytes('\t')
	time.Sleep(1e8)
}

func Peek(reader *bufio.Reader) {
	line, _ := reader.Peek(14)
	fmt.Printf("%s\n", line)
	//time.Sleep(1)
	fmt.Printf("%s\n", line)
}

func main() {
	samplePeek()
}
