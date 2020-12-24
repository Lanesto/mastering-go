package main

import "fmt"

func main() {
	const r1 = '€'
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(hex) r1: %x\n", r1)
	fmt.Printf("(string) r1: %s\n", r1)
	fmt.Printf("(char) r1: %c\n", r1)
	fmt.Println("A string is a collection of runes:", []byte("Mihalis"))
	aString := []byte("Mihalis")
	for i, r := range aString {
		fmt.Println(i, r)
		fmt.Printf("char: %c\n", aString[i])
	}
	fmt.Printf("%s\n", aString)
}
