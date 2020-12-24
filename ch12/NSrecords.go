package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Need a domain name")
		return
	}

	domain := args[1]
	nss, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ns := range nss {
		fmt.Println(ns.Host)
	}
}
