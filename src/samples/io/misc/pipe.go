package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")
	proverbs.WriteString("你好\n")

	piper, pipew := io.Pipe()

	go func() {
		defer pipew.Close()
		io.Copy(pipew, proverbs)
	}()

	io.Copy(os.Stdout, piper)
	piper.Close()
}
