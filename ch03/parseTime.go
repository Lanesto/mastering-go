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

	t, err := time.Parse("15:04", args[1])
	if err == nil {
		fmt.Println("Full:", t)
		fmt.Println("Time:", t.Hour(), t.Minute())
	} else {
		fmt.Println(err)
	}
}
