package main

import "fmt"

func hello(name string) {
	fmt.Println("hello", name)
}

type Hello func(string)

func (h Hello) Echo(name string) Hello {
	fmt.Println("echo", name)
	return h
}

func main() {
	var f Hello
	f = hello
	f.Echo("foo")("world")

	f1 := Hello(hello)
	f1.Echo("bar")
}
