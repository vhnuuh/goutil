//http://c.biancheng.net/view/5569.html
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("go语言练习")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	n, err = r.Read(buf[:])
	if err != nil {
		fmt.Println("read error", err)
	}
}
