package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var timeout = time.Duration(time.Second)

func withTimeout(network, host string) (net.Conn, error) {
	// conn, err := net.DialTimeout(network, host, timeout * time.Second)
	conn, err := net.Dial(network, host)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s URL TIMEOUT\n", filepath.Base(os.Args[0]))
		return
	}

	if len(os.Args) == 3 {
		temp, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Using default timeout")
		} else {
			timeout = time.Duration(time.Duration(temp) * time.Second)
			fmt.Println("Timeout set to", timeout)
		}
	}

	URL := os.Args[1]
	t := http.Transport{
		Dial: withTimeout,
	}
	client := http.Client{
		Transport: &t,
	}
	data, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
		}
	}
}
