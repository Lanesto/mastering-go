package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: %s [numeric [numeric [numeric ...]]]\n", args[0])
		return
	}

	args = args[1:]
	var sum float64
	n := 0
	for i, arg := range args {
		f, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			sum += f
			n++
		} else {
			fmt.Printf("warning: given %d-th argument '%s' is not numeric\n", i, arg)
		}
	}
	if n == 0 {
		fmt.Println("error: none of given arguments are numeric")
		os.Exit(1)
	}

	avg := sum / float64(n)
	fmt.Println(avg)
}
