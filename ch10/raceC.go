package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me a natural number!")
		os.Exit(1)
	}

	numGR, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup
	var i int
	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			k[i] = i
		}(i)
	}

	k[2] = 10
	waitGroup.Wait()
	fmt.Printf("k = %#v\n", k)
}
