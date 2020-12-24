package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: %s column <file1> [<file2> [....<fileN]]\n", filepath.Base(args[0]))
		return
	}

	column, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("column value is not an integer:", column)
		os.Exit(1)
	} else if column < 0 {
		fmt.Println("Invalid column number!")
		os.Exit(1)
	}

	for _, filename := range args[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file: %s\n", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file: %s\n", err)
			}
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println((data[column-1]))
			}
		}
	}
}
