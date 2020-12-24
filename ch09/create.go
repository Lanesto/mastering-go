package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 10, "Number of goroutines")
	flag.Parse()
	fmt.Printf("Going to create %d goroutines.\n", n)
	for i := 0; i < n; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("\nExiting")
}
