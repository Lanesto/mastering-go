package main

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
)

func sqrt(f float64, step uint) *big.Float {
	k := 0
	result := new(big.Float).SetPrec(step).SetFloat64(float64(0.5) * (float64(1.0) + (float64(f) / float64(1.0))))
	for {
		if k > int(step) {
			break
		}
		temp := new(big.Float).SetFloat64(f)
		temp.Quo(temp, result)
		result.Add(result, temp)
		result.Mul(result, new(big.Float).SetFloat64(0.5))
		k++
	}
	return result
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("Usage: %s <number> <step>\n", filepath.Base(args[0]))
		return
	}

	f, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("number must be floating point number")
		os.Exit(1)
	}

	step, err := strconv.ParseInt(args[2], 10, 32)
	if err != nil {
		fmt.Println("step must be integer")
		os.Exit(1)
	}

	fmt.Println(sqrt(f, uint(step)))
}
