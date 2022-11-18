//http://c.biancheng.net/view/5569.html
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("C语言中文网,Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadBytes(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadBytes(delim)
	fmt.Println(string(line), err)
}
