package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	var str string
	if len(args) != 2 {
		str = fmt.Sprintf("Usage: %s string", args[0])
	} else {
		str = args[1]
	}

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, str)
	io.WriteString(os.Stderr, "\n")
}
