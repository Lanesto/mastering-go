package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var password = secret{password: "myPassword"}

type secret struct {
	rwm      sync.RWMutex
	m        sync.Mutex
	password string
}

func change(c *secret, password string) {
	c.rwm.Lock()
	fmt.Println("LChange")
	time.Sleep(3 * time.Second)
	c.password = password
	c.rwm.Unlock()
}

func show(c *secret) string {
	c.rwm.RLock()
	fmt.Print("show")
	time.Sleep(3 * time.Second)
	defer c.rwm.RUnlock()
	return c.password
}

func showWithLock(c *secret) string {
	c.m.Lock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	defer c.m.Unlock()
	return c.password
}

func main() {
	var showFunction = func(c *secret) string { return "" }
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex")
		showFunction = show
	} else {
		fmt.Println("Using sync.Mutex")
		showFunction = showWithLock
	}

	var waitGroup sync.WaitGroup
	fmt.Println("Pass:", showFunction(&password))

	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Go pass:", showFunction(&password))
		}()
	}

	go func() {
		waitGroup.Add(1)
		defer waitGroup.Done()
		change(&password, "123456")
	}()

	waitGroup.Wait()
	fmt.Println("pass:", showFunction(&password))
}
