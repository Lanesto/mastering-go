package main

import (
	"fmt"
	"net"
	"os"
)

func lookupIP(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func lookupHostname(hostname string) ([]string, error) {
	ips, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide an argument")
		return
	}

	input := args[1]
	ip := net.ParseIP(input)
	if ip == nil {
		ips, err := lookupHostname(input)
		if err == nil {
			for _, ip := range ips {
				fmt.Println(ip)
			}
		}
	} else {
		hosts, err := lookupIP(input)
		if err == nil {
			for _, hostname := range hosts {
				fmt.Println(hostname)
			}
		}
	}
}
