//http://c.biancheng.net/view/5132.html
package main

import (
	"fmt"
	"github.com/codegangsta/inject"
	"reflect"
)

type SpecialString interface{}

func main() {
	inj := inject.New()
	inj.Map("C语言中文网")
	inj.MapTo("Golang", (*SpecialString)(nil))
	inj.Map(20)
	fmt.Println("string is valid?", inj.Get(reflect.TypeOf("Go语言入门教程")).
		IsValid())
	fmt.Println("special string is valid", inj.Get(inject.InterfaceOf((*SpecialString)(nil))).IsValid())
	fmt.Println("int is valid?", inj.Get(reflect.TypeOf(18)).IsValid())
	fmt.Println("[]byte is valid?", inj.Get(reflect.TypeOf([]byte("Golang"))).
		IsValid())
	inj2 := inject.New()
	inj2.Map([]byte("test"))
	inj.SetParent(inj2)
	fmt.Println("[]byte is valid?", inj.Get(reflect.TypeOf([]byte("Golang"))).
		IsValid())
}
