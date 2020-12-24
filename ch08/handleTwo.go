package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("handleSignal() caught:", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGUSR1)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
			case syscall.SIGUSR1:
				handleSignal(sig)
				return
			}
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(1 * time.Second)
	}
}
