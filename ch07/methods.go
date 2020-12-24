package main

import "fmt"

type twoInts struct {
	X int64
	Y int
}

func regularFunction(a, b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

func (a twoInts) method(b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

func main() {
	i := twoInts{1, 2}
	j := twoInts{-5, -2}
	fmt.Println(regularFunction(i, j))
	fmt.Println(i.method(j))
}
