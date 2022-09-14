package main

import (
	"fmt"
	"reflect"
)

// User 声明结构体
type User struct {
	Id   int
	Name string `json:"name" id:"100"`
}

func main() {
	//创建结构体实例
	ins := User{Id: 1, Name: "root"}
	//获取结构体实例的反射类型对象
	typ := reflect.TypeOf(ins)
	//通过字段名获取信息
	field, ok := typ.FieldByName("Name")
	if ok {
		tag := field.Tag
		tagName := tag.Get("json")
		tagId := tag.Get("id")
		fmt.Printf("name = %v, id = %v\n", tagName, tagId) //name = name, id = 100
	}
}
