package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1
	fmt.Println(areaIntf.Area())

	if sv, ok := sq1.(Shaper); ok {
		fmt.Printf("v implements String(): %s\n", sv)
	}

	//switch t := areaIntf.(*Square) {
	//case *Square:
	//	fmt.Printf("Type Square %T with value %v\n", t, t)
	//case *Circle:
	//	fmt.Printf("Type Circle %T with value %v\n", t, t)
	//case nil:
	//	fmt.Printf("nil value: nothing to check?\n")
	//default:
	//	fmt.Printf("Unexpected type %T\n", t)
	//}
}
