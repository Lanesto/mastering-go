package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var cache = map[int]int{
	0: 1,
	1: 1,
}

func f(n int) int {
	if n <= 0 {
		return -1
	} else if n <= 2 {
		return 1
	}

	var ret int
	if ret, ok := cache[n]; ok {
		return ret
	}

	ret = f(n-2) + f(n-1)
	cache[n] = ret
	return ret
}

func handleConn(c net.Conn) {
	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}

		temp := strings.TrimSpace(string(data))
		if temp == "STOP" {
			break
		}

		fibo := "-1\n"
		n, err := strconv.Atoi(temp)
		if err == nil {
			fibo = strconv.Itoa(f(n)) + "\n"
		}
		c.Write([]byte(string(fibo)))
	}
	time.Sleep(5 * time.Second)
	c.Close()
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + args[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("Serving")
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConn(c)
	}
}
