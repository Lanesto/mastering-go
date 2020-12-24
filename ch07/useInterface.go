package main

import (
	"fmt"
	"math"
	"myInterface"
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return 4 * s.X
}

func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func Calculate(x myInterface.Shape) {
	switch x.(type) {
	case circle:
		fmt.Println("Is a circle")
	case square:
		fmt.Println("Is a square")
	default:
		fmt.Println("I don't know")
	}

	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

func main() {
	x := square{10}
	fmt.Println("Perimeter:", x.Perimeter())
	Calculate(x)

	y := circle{5}
	Calculate(y)
}
