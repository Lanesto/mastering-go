package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: %s integer [integer [integer ...]]\n", args[0])
		return
	}

	args = args[1:]
	var sum int64 = 0
	for i, arg := range args {
		// Automatically implies base of given integer with its prefix
		// 0b for binary, 0o and 0 for octal, 0x for hexadecimal
		n, err := strconv.ParseInt(arg, 0, 64)
		if err == nil {
			sum += n
		} else {
			fmt.Printf("warning: given %d-th argument '%s' is not numeric\n", i, arg)
		}
	}
	fmt.Println(sum)
}
