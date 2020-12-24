package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generate(start int64, length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = byte((start + rand.Int63n(94)) % 128)
	}
	return string(result)
}

func main() {
	args := os.Args
	length, genCount, choice := 28, 3, -1
	var err error

	argCount := len(args)
	switch {
	case argCount > 3:
		choice, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Error while converting <choice> to numeric:", err)
			os.Exit(1)
		} else if choice <= 0 {
			fmt.Println("Error: <choice> must be positive integer")
		}
		fallthrough
	case argCount > 2:
		genCount, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error while converting <genCount> to numeric:", err)
			os.Exit(1)
		} else if genCount <= 0 {
			fmt.Println("Error: <genCount> must be positive integer")
		} else if genCount == 1 {
			fmt.Println("Info: <choice> set to 1 (<choice> <= <genCount>)")
			choice = 1
		}
		fallthrough
	case argCount > 1:
		length, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error while converting <length> to numeric:", err)
			os.Exit(1)
		} else if length <= 0 {
			fmt.Println("Error: <length> must be positive integer")
		}
	default:
		fmt.Printf("Using default values; usage: %s <length> <genCount> <choice> \n", os.Args[0])
	}

	rand.Seed(time.Now().Unix())
	candidates := make([]string, genCount)
	fmt.Println("Generating random passwords:")
	for i := 0; i < genCount; i++ {
		candidates[i] = generate('!', length)
		fmt.Printf("\tCandidate %d: %s\n", i+1, candidates[i])
	}
	fmt.Println()
	if choice == -1 {
		fmt.Printf("Please select one between %d-%d: ", 1, genCount)
		for {
			fmt.Scanf("%d", &choice)
			if choice <= 0 || choice > genCount {
				fmt.Printf("Invalid choice - try again please: ")
			} else {
				break
			}
		}
	}
	fmt.Println(candidates[choice-1])
}
