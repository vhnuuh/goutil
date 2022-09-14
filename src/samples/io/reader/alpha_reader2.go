package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type alphaReader struct {
	// alphaReader 里组合了标准库的 io.Reader
	reader io.Reader
}

func newAlphaReader(reader io.Reader) *alphaReader {
	return &alphaReader{reader: reader}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}

func print(r io.Reader) {
	p := make([]byte, 4)
	for {
		n, err := r.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}

func main() {
	reader := newAlphaReader(strings.NewReader("Hello! It's 9am, where is the sun?"))
	print(reader)

	file, err := os.Open("./alpha_reader.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader = newAlphaReader(file)
	print(reader)
}
