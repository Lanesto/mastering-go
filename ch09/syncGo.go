package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 20, "Number of goroutines")
	flag.Parse()
	fmt.Printf("Going to create %d goroutines\n", *n)

	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Add(1)
	for i := 0; i < *n; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Printf("%#v\n", waitGroup)
	fmt.Println("\nExiting")
}
