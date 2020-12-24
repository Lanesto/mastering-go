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
		str = fmt.Sprintf("%s %s", args[0], args[1])
	}

	io.WriteString(os.Stdout, str)
	io.WriteString(os.Stdout, "\n")
}
