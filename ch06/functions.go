package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func doubleSquare(x int) (int, int) {
	return 2 * x, x * x
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s <number>\n", filepath.Base(args[0]))
		return
	}

	y, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	square := func(s int) int {
		return s * s
	}
	fmt.Println("The square of", y, "is", square(y))
	double := func(s int) int {
		return s + s
	}
	fmt.Println("The double of", y, "is", double(y))

	fmt.Println(doubleSquare(y))
	d, s := doubleSquare(y)
	fmt.Println(d, s)
}
