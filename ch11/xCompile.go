package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("On a", runtime.GOARCH, "machine")
	fmt.Println("with Go version", runtime.Version())
}
