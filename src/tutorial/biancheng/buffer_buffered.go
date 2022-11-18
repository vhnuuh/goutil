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
	var buf [14]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	rn := r.Buffered()
	fmt.Println(rn)
	n, err = r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	rn = r.Buffered()
	fmt.Println(rn)
}
