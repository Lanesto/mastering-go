package main

import "fmt"

const (
	sq4_0 = 1 << (iota * 2)
	sq4_1
	sq4_2
	sq4_3
	sq4_4
	sq4_5
)

const (
	mon = iota
	tue
	wed
	thu
	fri
	sat
	sun
)

func main() {
	fmt.Println("4^0:", sq4_0)
	fmt.Println("4^1:", sq4_1)
	fmt.Println("4^2:", sq4_2)
	fmt.Println("4^3:", sq4_3)
	fmt.Println("4^4:", sq4_4)
	fmt.Println("4^5:", sq4_5)

	fmt.Println("Mon:", mon)
	fmt.Println("Tue:", tue)
	fmt.Println("Wed:", wed)
	fmt.Println("Thu:", thu)
	fmt.Println("Fri:", fri)
	fmt.Println("Sat:", sat)
	fmt.Println("Sun:", sun)
}
