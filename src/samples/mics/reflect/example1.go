package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name    string
	Address string
}

func main() {
	content := "Name"
	p := new(People)
	//
	v := reflect.ValueOf(p).Elem()
	fmt.Println(v.FieldByName(content).CanSet())

	v.FieldByName(content).SetString("jiangf")
	v.FieldByName("Address").SetString("GuangZhou")
	fmt.Println(p)

}
