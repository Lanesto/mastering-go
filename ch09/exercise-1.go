package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	// "path/filepath"
	"sync"
)

func countAll(filename, substr string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		text := scanner.Text()
		count += strings.Count(text, substr)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func main() {
	var substr string
	var n int
	flag.StringVar(&substr, "s", "", "substr to count")
	flag.IntVar(&n, "n", 1, "number of goroutines")
	flag.Parse()
	if len(flag.Args()) == 0 {
		// fmt.Printf("Usage: %s [options] <file1> [<file2> [...]]\n", filepath.Base(os.Args[0]))
		flag.Usage()
		return
	}

	var sem sync.WaitGroup
	total := 0
	for i, filename := range flag.Args() {
		label := fmt.Sprintf("goroutine %d", i)
		sem.Add(1)
		go func(label, filename, substr string) {
			defer sem.Done()
			n, err := countAll(filename, substr)
			if err != nil {
				fmt.Println(label, err)
			} else {
				total += n
			}
		}(label, filename, substr)
	}
	sem.Wait()
	fmt.Println(total)
}
