package main

import (
	"fmt"
	"net"
)

func main() {
	ifs, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range ifs {
		fmt.Printf("Interface: %v\n", i.Name)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}

		addrs, err := byName.Addrs()
		for k, v := range addrs {
			fmt.Printf("Interface address %#v: %v\n", k, v.String())
		}
		fmt.Println()
	}
}
