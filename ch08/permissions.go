package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("Usage: %s <file1> [<file2> [<file3> ...]]\n", filepath.Base(os.Args[0]))
		return
	}

	for _, filename := range flag.Args() {
		info, _ := os.Stat(filename)
		mode := info.Mode()
		fmt.Println(filename, "mode is", mode.String()[1:10])
	}
}
