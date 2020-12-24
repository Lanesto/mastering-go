package main

import (
	"fmt"
	"time"
)

func fnA(a, b chan struct{}) {
	<-a
	fmt.Println("fnA()!")
	time.Sleep(time.Second)
	close(b)
}

func fnB(a, b chan struct{}) {
	<-a
	fmt.Println("fnB()!")
	close(b)
}

func fnC(a chan struct{}) {
	<-a
	fmt.Println("fnC()!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go fnC(z)
	go fnA(x, y)
	go fnC(z)
	go fnB(y, z)
	go fnC(z)

	close(x)
	time.Sleep(3 * time.Second)
}
