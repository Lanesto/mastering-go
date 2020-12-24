package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var bufferSize int
var fileSize int

func random(min, max int) int {
	return min + rand.Intn(max-min)
}

func createBuffer(buf *[]byte, count int) {
	*buf = make([]byte, count)
	if count == 0 {
		return
	}

	for i := 0; i < count; i++ {
		intByte := byte(random(0, 100))
		if len(*buf) > count {
			return
		}
		*buf = append(*buf, intByte)
	}
}

func create(dst string, b, f int) error {
	_, err := os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}
	buf := make([]byte, 0)
	for {
		createBuffer(&buf, b)
		buf = buf[:b]
		if _, err := destination.Write(buf); err != nil {
			return err
		}
		if f < 0 {
			break
		}
		f = f - len(buf)
	}
	return err
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need bufferSize fileSize!")
		return
	}
	output := "/tmp/randomFile"
	bufferSize, _ := strconv.Atoi(os.Args[1])
	fileSize, _ := strconv.Atoi(os.Args[2])
	err := create(output, bufferSize, fileSize)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Remove(output)
	if err != nil {
		fmt.Println(err)
	}
}
