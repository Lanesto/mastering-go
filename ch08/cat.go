package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func printFile(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, filename := range args[1:] {
		if err := printFile(filename); err != nil {
			fmt.Println(err)
		}
	}
}
