package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

const ipv4FieldPattern = `(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[0-9]?[0-9])`
const ipv4Pattern = ipv4FieldPattern + "\\." + ipv4FieldPattern + "\\." + ipv4FieldPattern + "\\." + ipv4FieldPattern

var ipv4 = regexp.MustCompile(ipv4Pattern)
var counter = make(map[string]int)

func countIPv4All(line string) int {
	match := ipv4.FindAllString(line, -1)
	for _, address := range match {
		trial := net.ParseIP(address)
		if trial.To4() == nil {
			fmt.Printf("found malformed IPv4 address: %s\n", address)
			continue
		} else {
			counter[address]++
		}
	}
	return len(match)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: %s <file1> [<file2> [<file3> ...]]\n", filepath.Base(args[0]))
		return
	}

	for _, filename := range args[1:] {
		f, err := os.Open(args[1])
		if err != nil {
			fmt.Printf("could not open file %q; skipping\n", filename)
			continue
		}
		defer f.Close()
		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error occured while reading file %s; error=%s\n", filename, err)
				break
			}

			countIPv4All(line)
		}
	}

	maxCount := -1
	mostFrequentAddress := ""
	for address, count := range counter {
		if count > maxCount {
			maxCount = count
			mostFrequentAddress = address
		}
	}
	fmt.Printf("found most frequent IPv4 address for all given %d files with %d occurences: %s\n", len(args[1:]), maxCount, mostFrequentAddress)
	fmt.Println("\n* hint: to test it, use 'grep -o $address $filename | wc -l'")
}
