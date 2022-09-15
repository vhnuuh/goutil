package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	OrderId    int
	customerId int
}

type Employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		typ := reflect.TypeOf(q)
		val := reflect.ValueOf(q)
		for i := 0; i < typ.NumField(); i++ {
			//fmt.Println(val.Elem().Field(i).Interface())
			switch val.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, val.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, val.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, val.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

func main() {
	o := Order{
		OrderId:    456,
		customerId: 56,
	}

	fmt.Println(reflect.ValueOf(&o).Kind())
	fmt.Println(reflect.ValueOf(&o).Elem().Kind())
	fmt.Println(reflect.ValueOf(&o).Elem().Field(0).Interface())
	fmt.Println(reflect.ValueOf(o).Kind())
	fmt.Println(reflect.ValueOf(o).Elem().Kind())
	fmt.Println(reflect.ValueOf(o).Elem().Field(0).Interface())
	createQuery(o)

	e := Employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)

}
