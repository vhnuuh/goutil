//http://c.biancheng.net/view/5569.html
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	bl, err := r.Peek(8)
	fmt.Println(string(bl), err)
	bl, err = r.Peek(14)
	fmt.Println(string(bl), err)
	bl, err = r.Peek(20)
	fmt.Println(string(bl), err)
}
