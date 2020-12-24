package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s string\n", filepath.Base(args[0]))
		return
	}

	d, err := time.Parse("02 January 2006", args[1])
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	} else {
		fmt.Println(err)
	}
}
