package main

import "fmt"

func main() {
	var grade = "B"
	var marks = 0

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "y"
	}

	switch {
	case grade == "A":
		fmt.Printf("A\n")
	case grade == "B", grade == "C":
		fmt.Printf("B C\n")
	case grade == "D":
		fmt.Printf("D\n")
	case grade == "F":
		fmt.Printf("F\n")
	default:
		fmt.Printf("error\n")
		fmt.Printf("error1\n")
	}
	fmt.Printf("your grade is %s\n", grade)
}
