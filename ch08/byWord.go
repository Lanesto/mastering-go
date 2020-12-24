package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "io"
	"os"
	"path/filepath"
	"regexp"
)

var wordRegexp = regexp.MustCompile(`[^\s]+`)

func wordByWord(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	/*
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}

			words := wordRegexp.FindAllString(line, -1)
			for _, w := range words {
				fmt.Println(w)
			}
		}
	*/
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("Usage: %s <file1> [<file2> ...]\n", filepath.Base(os.Args[0]))
		return
	}

	for _, file := range flag.Args() {
		err := wordByWord(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
