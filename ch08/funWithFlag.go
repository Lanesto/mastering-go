package main

import (
	"flag"
	"fmt"
	"strings"
)

type namesFlag struct {
	Names []string
}

func (s *namesFlag) GetNames() []string {
	return s.Names
}

func (s *namesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *namesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("Cannot use names flag more than once")
	}

	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames namesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Mihalis", "The name")
	flag.Var(&manyNames, "names", "Comma-separated list")
	flag.Parse()

	fmt.Println("-k:", *minusK)
	fmt.Println("-o:", *minusO)
	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("remaining command-line arguments:")
	for index, value := range flag.Args() {
		fmt.Println(index, ":", value)
	}
}
