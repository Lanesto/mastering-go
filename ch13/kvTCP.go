package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const dataFile = "dataFile.gob"

var data = make(map[string]record)

const welcomeMsg = "Welcome to the key-value store\n"

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

func show(c net.Conn) {
	for k, v := range data {
		resp := fmt.Sprintf("%s: %v\n", k, v)
		c.Write([]byte(resp))
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

func handleConn(c net.Conn) {
	c.Write([]byte(welcomeMsg))
	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		text := strings.TrimSpace(data)
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
		var respMsg string
		switch command {
		case "show":
			show(c)
		case "stop":
			c.Close()
			return
		case "remove":
			if !remove(key) {
				respMsg = "Operation failed: remove\n"
			} else {
				respMsg = "Operation succeed: remove\n"
			}
		case "insert":
			r := record{v1, v2, v3}
			if !insert(key, r) {
				respMsg = "Operation failed: insert\n"
			} else {
				respMsg = "Operation succeed: insert\n"
			}
		case "get":
			r := get(key)
			if r != nil {
				respMsg = fmt.Sprintf("%v\n", *r)
			} else {
				respMsg = "Did not find key\n"
			}
		case "update":
			r := record{v1, v2, v3}
			if !update(key, r) {
				respMsg = "Operation failed: update\n"
			} else {
				respMsg = "Operation succeed: update\n"
			}
		default:
			respMsg = fmt.Sprintf("Unknown command: '%s' please try again\n", command)
		}
		c.Write([]byte(respMsg))
		if err := save(); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a port number")
		return
	}

	PORT := ":" + args[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	if err := load(); err != nil {
		log.Println(err)
	}
	defer func() {
		if err := save(); err != nil {
			log.Println(err)
		}
	}()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		go handleConn(c)
	}
}
