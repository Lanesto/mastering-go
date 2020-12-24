package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 3 {
		fmt.Printf("Usage: %s <file> <src> <dest>\n", filepath.Base(os.Args[0]))
		return
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	src := flag.Arg(1)
	dst := flag.Arg(2)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, src, dst, -1)
		fmt.Println(line)
	}
}
