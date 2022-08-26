package main

func If(condition bool, trueval, falseValue interface{}) interface{} {
	if condition {
		return trueval
	}
	return falseValue
}

func main() {
	a, b := 2, 3
	max := If(a > b, a, b).(int)
	print(max)
}
