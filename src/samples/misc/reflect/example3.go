// https://blog.csdn.net/qq_44470091/article/details/110655177

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name  string `json:"name"`
	Count int
}

func (p *Person) Print() {
	fmt.Println("print:", p)
}

func (p *Person) CountAdd(num int) int {
	return p.Count + num
}

func main() {
	test(&Person{
		Name:  "lei",
		Count: 2,
	})
}

func test(body interface{}) {
	//TypeOf会返回目标数据的类型，比如int/float/struct/指针等
	typ := reflect.TypeOf(body)
	//ValueOf返回目标数据的的值
	val := reflect.ValueOf(body)
	if val.Elem().Kind() != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	fmt.Println(typ)
	fmt.Println(val)
	for i := 0; i < val.Elem().NumField(); i++ {
		field := typ.Elem().Field(i)
		value := val.Elem().Field(i)
		fmt.Println("type1:", field)
		fmt.Println("value1", value)
		switch value.Kind() {
		case reflect.Int:
			value.SetInt(88)
		case reflect.String:
			value.SetString("Test")
		default:
			fmt.Println("类型不支持")
		}
		fmt.Println("type2:", field)
		fmt.Println("value2:", value)
		fmt.Println(field.Tag.Get("json"))
	}
	//   除解析一个接口的结构字段和方法外，还可以对注册在结构上的方法进行调用
	//   参数和返回值都是reflect包中Value型的切片，需要经过转换
	call := val.Method(0).Call([]reflect.Value{reflect.ValueOf(2)})
	fmt.Println("返回值：", call[0])
	val.MethodByName("Print").Call(nil)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		fmt.Println(method.Name, method.Type)
		if method.Name == "CountAdd" {
			retInfo := method.Func.Call([]reflect.Value{val, reflect.ValueOf(2)})
			fmt.Println("返回值：", retInfo[0])
		}
	}
	method, ok := typ.MethodByName("Print")
	if ok {
		method.Func.Call([]reflect.Value{val})
	}
}
