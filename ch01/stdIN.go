package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println("> ", scanner.Text())
	}
}
