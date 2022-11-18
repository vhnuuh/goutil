package main

import "fmt"

type Phone interface {
	call(message string)
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call(message string) {
	fmt.Println("call from nokia", message)
}

type IPhone struct {
}

func (iphone IPhone) call(message string) {
	fmt.Println("call from iphone", message)
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call("hello")

	phone = new(IPhone)
	phone.call("hello")
}
