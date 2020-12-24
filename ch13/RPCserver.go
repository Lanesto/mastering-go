package main

import (
	"fmt"
	"math"
	"net"
	"net/rpc"
	"os"
	"sharedrpc"
)

// MyInterface is implementation of sharedrpc.MyInterface
type MyInterface struct{}

// Power returns x^y
func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

// Multiply returns multiplication of arguments
func (t *MyInterface) Multiply(arguments *sharedrpc.MyFloats, reply *float64) error {
	*reply = arguments.A1 * arguments.A2
	return nil
}

// Power returns power
func (t *MyInterface) Power(arguments *sharedrpc.MyFloats, reply *float64) error {
	*reply = Power(arguments.A1, arguments.A2)
	return nil
}

func main() {
	PORT := ":1234"
	args := os.Args
	if len(args) != 1 {
		PORT = ":" + args[1]
	}

	MyInterface := new(MyInterface)
	rpc.Register(MyInterface)
	t, err := net.ResolveTCPAddr("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}
}
