//type assertions
package main

import "fmt"

func main() {
	var a interface{} = 100

	if aa, ok := a.(int); ok {
		fmt.Println(aa)
	}

	// 如果类型不同，会抛出panic
	ab := a.(string)
	fmt.Println(ab)
}
