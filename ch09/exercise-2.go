package main

import (
	"fmt"
	// "os"
	"flag"
)

// Each function should close 'out' channel to send signal "i'm done"
type pipeFn func(in <-chan int, out chan<- int)

type pipeSpec struct {
	Pipe pipeFn
	Num  int
}

func connect(recipes ...pipeSpec) (in, out chan int) {
	in, out = make(chan int), make(chan int)

	// First pipe to input
	before := in
	numRecipes := len(recipes)
	for i, recipe := range recipes {
		var tunnel chan int
		// Last pipe to output
		if i == (numRecipes - 1) {
			tunnel = out
		} else {
			tunnel = make(chan int)
		}
		// Run recipe.Num goroutines
		for j := 0; j < recipe.Num; j++ {
			go recipe.Pipe(before, tunnel)
		}
		before = tunnel
	}
	return in, out
}

func squarer(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}

func totaler(in <-chan int, out chan<- int) {
	sum := 0
	for n := range in {
		sum += n
	}
	out <- sum
}

func main() {
	min := flag.Int("min", 1, "minimal number of range")
	max := flag.Int("max", 10, "maximum number of range")
	n := flag.Int("n", 1, "number of squarer pipes") // Causes error! for experiment only(use 1 for safe)
	flag.Parse()

	in, out := connect(pipeSpec{squarer, *n}, pipeSpec{totaler, 1})
	for i := *min; i <= *max; i++ {
		in <- i
	}
	close(in)
	fmt.Println(<-out)
}
