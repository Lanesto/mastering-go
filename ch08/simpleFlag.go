package main

import (
	"flag"
	"fmt"
)

func main() {
	minusK := flag.Bool("k", true, "k")
	minus0 := flag.Int("0", 1, "0")
	flag.Parse()

	valueK := *minusK
	value0 := *minus0
	value0++
	fmt.Println("-k:", valueK)
	fmt.Println("-0:", value0)
}
