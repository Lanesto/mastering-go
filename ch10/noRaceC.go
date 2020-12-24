package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var m sync.Mutex

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
			m.Lock()
			k[i] = i
			m.Unlock()
		}(i)
	}

	waitGroup.Wait()
	k[2] = 10
	fmt.Printf("k = %#v\n", k)
}
