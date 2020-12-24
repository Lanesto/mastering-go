package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func charByChar(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		for _, c := range line {
			fmt.Println(string(c))
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("Usage: %s <file1> [<file2> ...]\n", filepath.Base(os.Args[0]))
		return
	}

	for _, file := range flag.Args() {
		err := charByChar(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
