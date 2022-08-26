package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(int(time.Now().Weekday()))
	fmt.Println(int(time.Now().Month()))
	var a float64
	a = 3.1
	b := int(a)
	fmt.Println(b)
}
