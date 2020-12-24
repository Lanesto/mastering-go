package main

import "C"

import "fmt"

// PrintMessage prints static message to stdout
//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

// Multiply returns multiplication of two given integers
//export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {
}
