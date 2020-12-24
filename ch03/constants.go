package main

import "fmt"

type digit int
type power2 int

const pi = 3.1415926

const (
	c1 = "c1c1c1"
	c2 = "c2c2c2"
	c3 = "c3c3c3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(pi)

	const (
		Zero digit = iota
		One
		Two
		Three
		Four
	)
	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 power2 = 1 << iota
		_
		p2_2
		_
		p2_4
		_
		p2_6
	)
	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
}
