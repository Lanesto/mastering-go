package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type client struct {
	id      int
	integer int
}

type dat struct {
	job    client
	square int
}

var (
	size    = 2
	clients = make(chan client, size)
	data    = make(chan dat, size)
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		fmt.Println("worker received:", c.integer)
		square := c.integer * c.integer
		output := dat{c, square}
		fmt.Println("worker sending:", output)
		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}

func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers")
		os.Exit(1)
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	go create(nJobs)
	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint:", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()

	makeWP(nWorkers)
	fmt.Printf(": %v\n", <-finished)
}
