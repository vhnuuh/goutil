//https://studygolang.com/articles/13178
package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func sample(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	sample(o)
}
