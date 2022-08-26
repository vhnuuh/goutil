package main

import "fmt"

func main() {
	var t interface{} = 100
	switch i := t.(type) {
	case float32:
		fmt.Printf("type: %T, value %v\n", i, i)
	case float64:
		fmt.Printf("type: %T, value %v\n", i, i)
	case int:
		fmt.Printf("type: %T, value %v\n", i, i)
	case string:
		fmt.Printf("type: %T, value %v\n", i, i)
	default:
		fmt.Println("other type")
	}
}
