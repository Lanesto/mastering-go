package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var closeChannel = false

var data = make(map[int]bool)

func random(min, max int) int {
	return min + rand.Intn(max-min)
}

func first(min, max int, out chan<- int) {
	for {
		if closeChannel {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		_, ok := data[x]
		if ok {
			closeChannel = true
		} else {
			data[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	sum := 0
	count := 0
	for x := range in {
		sum += x
		count++
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
	fmt.Printf("The count of generated random numbers is %d\n", count)
}

func main() {
	min := flag.Int("min", 0, "Minimum of generated random number")
	max := flag.Int("max", 10, "Maximum of generated random number")
	flag.Parse()

	if *min > *max {
		fmt.Println("Swapping min, max becuz min > max")
		*min, *max = *max, *min
	}

	rand.Seed(time.Now().UnixNano())

	in := make(chan int)
	out := make(chan int)

	go first(*min, *max, in)
	go second(out, in)
	third(out)
}
