package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"
)

const dataFile = "dataFile.gob"

var data = make(map[string]record)

type record struct {
	Name    string
	Surname string
	ID      string
}

func save() error {
	fmt.Printf("Saving data to %s\n", dataFile)
	err := os.Remove(dataFile)
	if err != nil {
		log.Println("Could not remove existing data file:", err)
	}

	f, err := os.Create(dataFile)
	if err != nil {
		log.Println("Could not create data file:", err)
		return err
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)
	if err := encoder.Encode(data); err != nil {
		log.Println("Could not save data to", dataFile)
		return err
	}

	return nil
}

func load() error {
	fmt.Printf("Loading data from %s\n", dataFile)
	f, err := os.Open(dataFile)
	if err != nil {
		log.Println("Could not open data file:", err)
		return err
	}
	defer f.Close()

	decoder := gob.NewDecoder(f)
	if err := decoder.Decode(&data); err != nil {
		log.Println("Could not load data from", dataFile)
		return err
	}

	return nil
}

func insert(key string, r record) bool {
	if key == "" {
		log.Println("Empty key given; ignored")
		return false
	}

	if get(key) == nil {
		log.Printf("Added new item with key %s, value %s\n", key, r)
		data[key] = r
		return true
	}

	log.Printf("Key %s already occupied", key)
	return false
}

func remove(key string) bool {
	if get(key) != nil {
		delete(data, key)
		log.Printf("Removed item with key %s\n", key)
		return true
	}

	log.Printf("Tried to remove item with key %s but has no data to delete\n", key)
	return false
}

func get(key string) *record {
	log.Printf("Someone looked up key %s\n", key)
	if _, ok := data[key]; ok {
		r := data[key]
		return &r
	}

	return nil
}

func update(key string, r record) bool {
	data[key] = r
	log.Printf("Item with key %s updated to %s\n", key, r)
	return true
}

func show() {
	for k, v := range data {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func init() {
	fmt.Println("Initializing")
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could not open log file; %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
}

func main() {
	if err := load(); err != nil {
		log.Println(err)
	}
	defer func() {
		if err := save(); err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("Interaction: <command> <key> [<value1> <value2> <value3>]")
	fmt.Println("<command>: show stop remove insert get update")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		tokens := strings.Fields(text)
		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			fallthrough
		case 2:
			tokens = append(tokens, "")
			fallthrough
		case 3:
			tokens = append(tokens, "")
			fallthrough
		case 4:
			tokens = append(tokens, "")
		}

		command, key, v1, v2, v3 := tokens[0], tokens[1], tokens[2], tokens[3], tokens[4]
		switch command {
		case "show":
			show()
		case "stop":
			return
		case "remove":
			if !remove(key) {
				fmt.Println("Operation failed: remove")
			}
		case "insert":
			r := record{v1, v2, v3}
			if !insert(key, r) {
				fmt.Println("Operation failed: insert")
			}
		case "get":
			r := get(key)
			if r != nil {
				fmt.Printf("%v\n", *r)
			}
		case "update":
			r := record{v1, v2, v3}
			if !update(key, r) {
				fmt.Println("Operation failed: update")
			}
		default:
			fmt.Printf("Unknown command: '%s' please try again\n", command)
		}
	}
}
