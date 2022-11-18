package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Printf("please input password: \n")
	fmt.Scan(&a)
	if a == 5211314 {
		fmt.Printf("please input password again")
		fmt.Scan(&b)
		if b == 5211314 {
			fmt.Printf("password correct, open the door")
		} else {
			fmt.Printf("alarm")
		}
	} else {
		fmt.Printf("alarm")
	}
}
