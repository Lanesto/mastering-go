package main

import "fmt"

type float float64

type node struct {
	Value float
	Next  *node
}

var size = 0
var queue = new(node)

func push(t *node, v float) bool {
	if queue == nil {
		queue = &node{v, nil}
		size++
		return true
	}

	t = &node{v, nil}
	t.Next = queue
	queue = t

	size++

	return true
}

func pop(t *node) (float, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		queue = nil
		size--
		return t.Value, true
	}

	tmp := t
	for t.Next != nil {
		tmp = t
		t = t.Next
	}

	v := tmp.Next.Value
	tmp.Next = nil

	size--
	return v, true
}

func traverse(t *node) {
	if t == nil {
		fmt.Println("Empty Queue!")
		return
	}

	for t != nil {
		fmt.Printf("%f -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	push(queue, 10)
	fmt.Println("Size:", size)
	traverse(queue)

	v, b := pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	for i := 0; i < 5; i++ {
		push(queue, float(i))
	}
	traverse(queue)
	fmt.Println("Size:", size)

	v, b = pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	v, b = pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	traverse(queue)
}
