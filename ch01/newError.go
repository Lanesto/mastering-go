package main

import (
	"errors"
	"fmt"
)

var errCustom = errors.New("error in returnError() function")

func returnError(a, b int) error {
	if a == b {
		return errCustom
	}
	return nil
}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "error in returnError() function" {
		fmt.Println("!!")
	}
}
