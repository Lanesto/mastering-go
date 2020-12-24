package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: switch number")
		os.Exit(1)
	}

	number, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("This value is not an integer:", number)
	} else {
		switch {
		case number < 0:
			fmt.Println("Less than zero!")
		case number > 0:
			fmt.Println("Greater than zero!")
		default:
			fmt.Println("Zero")
		}
	}

	asString := args[1]
	switch asString {
	case "5":
		fmt.Println("Five!")
	case "0":
		fmt.Println("Zero!")
	default:
		fmt.Println("Do not care!")
	}

	var neg = regexp.MustCompile(`-`)
	var float = regexp.MustCompile(`\d?\.\d`)
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]*`)
	switch {
	case neg.MatchString(asString):
		fmt.Println("Negative number")
	case float.MatchString(asString):
		fmt.Println("Floating point")
	case email.MatchString(asString):
		fmt.Println("It is an email!")
		fallthrough
	default:
		fmt.Println("Something else!")
	}

	var aType error = nil
	switch aType.(type) {
	case nil:
		fmt.Println("It is nil interface!")
	default:
		fmt.Println("Not nil interface!")
	}
}
