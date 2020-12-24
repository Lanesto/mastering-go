package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)
	n, err := f.Read(buffer)
	if err == io.EOF {
		return nil
	} else if err != nil {
		fmt.Println(err)
		return nil
	}
	return buffer[:n]
}

func main() {
	defer func() {
		if c := recover(); c != nil {
			fmt.Println(c)
			fmt.Printf("Usage: %s -n <int> -f <filename>\n", filepath.Base(os.Args[0]))
			return
		}
	}()

	n := flag.Int("n", 10, "Number of bytes to read")
	filename := flag.String("f", "", "Name of file to read")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for {
		buf := readSize(f, *n)
		if buf != nil {
			fmt.Print(string(buf))
		} else {
			break
		}
	}
}
