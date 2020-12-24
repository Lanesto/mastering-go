package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var f *os.File = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.EqualFold(text, "STOP") {
			break
		}

		n, err := strconv.ParseInt(text, 0, 64)
		if err != nil {
			fmt.Println("> Non-integer item given!")
		} else {
			fmt.Println(">", n)
		}
	}
	fmt.Println("Exiting")
}
