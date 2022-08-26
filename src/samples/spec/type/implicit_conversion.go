package main

import (
	"fmt"
	"reflect"
)

type Handler func()

func a() Handler {
	return func() {}
}

func main() {
	var i interface{} = main
	_, ok := i.(func())
	fmt.Println(ok)
	_, ok = i.(Handler)
	fmt.Println(ok)
	a := (*Handler)(nil)
	fmt.Println(reflect.TypeOf(main) == reflect.TypeOf(a).Elem())
}
