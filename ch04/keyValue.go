package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type myElement struct {
	Name    string
	Surname string
	ID      string
}

var data = make(map[string]myElement)

func add(k string, n myElement) bool {
	if k == "" {
		return false
	}

	if lookup(k) == nil {
		data[k] = n
		log.Printf("Added new item with key %s, value %s\n", k, n)
		return true
	}

	return false
}

func remove(k string) bool {
	if lookup(k) != nil {
		delete(data, k)
		log.Printf("Removed item with key %s\n", k)
		return true
	}
	return false
}

func lookup(k string) *myElement {
	_, ok := data[k]
	if ok {
		n := data[k]
		log.Printf("Someone looked up key %s\n", k)
		return &n
	}
	return nil
}

func change(k string, n myElement) bool {
	data[k] = n
	log.Printf("Item with key %s updated to %s\n", k, n)
	return true
}

func print() {
	for k, v := range data {
		fmt.Printf("Key: %s - Value: %s\n", k, v)
	}
}

func main() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could not open log file; err=%v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)
		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "print":
			print()
		case "stop":
			return
		case "remove":
			if !remove(tokens[1]) {
				fmt.Println("Remove operation failed!")
			}
		case "add":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !add(tokens[1], n) {
				fmt.Println("Add operation failed!")
			}
		case "lookup":
			n := lookup(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", *n)
			}
		case "change":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !change(tokens[1], n) {
				fmt.Println("Update operation failed!")
			}
		default:
			fmt.Println("Unknown command - please try again!")
		}
	}
}
