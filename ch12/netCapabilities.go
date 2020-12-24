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
		fmt.Printf("Name: %v\n", i.Name)
		fmt.Println("Interface Flags:", i.Flags.String())
		fmt.Println("Interface MTU:", i.MTU)
		fmt.Println("Interface hardware address:", i.HardwareAddr)
		fmt.Println()
	}
}
