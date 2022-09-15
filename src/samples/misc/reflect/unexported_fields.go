// https://zhuanlan.zhihu.com/p/148231342

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Example struct {
	a string
}

func GetStructPtrUnExportedField(source interface{}, fieldName string) reflect.Value {
	v := reflect.ValueOf(source).Elem().FieldByName(fieldName)
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

func SetStructPtrUnExportedStrField(source interface{}, fieldName string, fieldValue interface{}) (err error) {
	v := GetStructPtrUnExportedField(source, fieldName)
	rv := reflect.ValueOf(fieldValue)
	if v.Kind() != rv.Kind() {
		return fmt.Errorf("invalid kind: expected kind %v, got kind: %v", v.Kind(), rv.Kind())
	}

	v.Set(rv)
	return nil
}

func main() {
	e := Example{
		a: "example",
	}

	unexported := GetStructPtrUnExportedField(&e, "a")
	fmt.Println(unexported)

	_ = SetStructPtrUnExportedStrField(&e, "a", "Hello,world")
	fmt.Println(e)
}
