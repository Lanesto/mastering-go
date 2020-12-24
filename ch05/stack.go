package main

import "fmt"

type node struct {
	Value int
	Next  *node
}

var size = 0
var stack = new(node)

func push(v int) bool {
	if stack == nil {
		stack = &node{v, nil}
		size = 1
		return true
	}

	temp := &node{v, stack}
	stack = temp
	size++
	return true
}

func pop(t *node) (int, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		size = 0
		stack = nil
		return t.Value, true
	}

	stack = stack.Next
	size--
	return t.Value, true
}

func traverse(t *node) {
	if size == 0 {
		fmt.Println("Empty Stack!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	stack = nil
	v, b := pop(stack)
	if b {
		fmt.Print(v, " ")
	} else {
		fmt.Println("Pop() failed!")
	}

	push(100)
	traverse(stack)
	push(200)
	traverse(stack)

	for i := 0; i < 10; i++ {
		push(i)
	}

	for i := 0; i < 15; i++ {
		v, b := pop(stack)
		if b {
			fmt.Print(v, " ")
		} else {
			break
		}
	}
	fmt.Println()
	traverse(stack)
}
