package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Entry struct {
	Number int
	Square int
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s <template>\n", filepath.Base(args[0]))
		return
	}

	tpl := args[1]
	data := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}
	var entries []Entry
	for _, i := range data {
		if len(i) == 2 {
			entries = append(entries, Entry{Number: i[0], Square: i[1]})
		}
	}

	t := template.Must(template.ParseGlob(tpl))
	t.Execute(os.Stdout, entries)
}
