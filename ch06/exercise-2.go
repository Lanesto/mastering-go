package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func sortNumNamedReturn(a, b, c int) (x, y, z int) {
	if a > b {
		// ?, a, ?, b, ?
		if c > a {
			x, y, z = c, a, b
		} else if c > b {
			x, y, z = a, c, b
		} else {
			x, y, z = a, b, c
		}
	} else {
		// ?, b, ?, a, ?
		if c > b {
			x, y, z = c, b, a
		} else if c > a {
			x, y, z = b, c, a
		} else {
			x, y, z = b, a, c
		}
	}
	return
}

func sortNum(a, b, c int) (int, int, int) {
	if a > b {
		// ?, a, ?, b, ?
		if c > a {
			return c, a, b
		} else if c > b {
			return a, c, b
		} else {
			return a, b, c
		}
	} else {
		// ?, b, ?, a, ?
		if c > b {
			return c, b, a
		} else if c > a {
			return b, c, a
		} else {
			return b, a, c
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Printf("Usage: %s <number> <number> <number>\n", filepath.Base(args[0]))
		return
	}

	a, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(sortNum(a, b, c))
	fmt.Println(sortNumNamedReturn(a, b, c))
}
