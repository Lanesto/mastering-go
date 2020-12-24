package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var errNoNumericArgument = errors.New("none of given command-line arguments are numeric")

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: %s numeric [numeric [numeric ...]]\n", args[0])
		return
	}

	defer func() {
		if c := recover(); c != nil {
			fmt.Printf("error: %v\n", c)
			os.Exit(1)
		}
	}()

	args = args[1:]
	values := make([]float64, 0, len(args))
	for i, arg := range args {
		val, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			values = append(values, val)
		} else {
			fmt.Printf("warning: given %d-th argument '%s' is not numeric\n", i, arg)
		}
	}
	if len(values) == 0 {
		panic(errNoNumericArgument)
	}

	min, max := values[0], values[0]
	for _, val := range values {
		if val < min {
			min = val
		} else if val > max {
			max = val
		}
	}
	fmt.Println("min:", min)
	fmt.Println("max:", max)
}
