package main

import "fmt"

type node struct {
	Value int
	Next  *node
}

const minInt = -int(^uint(0)>>1) - 1

// Returns new HEAD pointer after inserting new element
func addNode(t *node, v int) *node {
	if t == nil {
		return &node{v, nil}
	}

	temp := &node{minInt, t}
	t = temp
	for t.Next != nil {
		if t.Next.Value == v {
			fmt.Println("Node already exists:", v)
			break
		} else if t.Next.Value > v {
			t.Next = &node{v, t.Next}
			break
		}
		t = t.Next
	}

	if t.Next == nil {
		t.Next = &node{v, nil}
	}

	return temp.Next
}

func traverse(t *node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func lookupNode(t *node, v int) bool {
	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func size(t *node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	var root *node
	root = addNode(root, 1)
	root = addNode(root, -1)
	traverse(root)
	root = addNode(root, 10)
	root = addNode(root, 5)
	root = addNode(root, 45)
	root = addNode(root, 5)
	root = addNode(root, 5)
	traverse(root)
	root = addNode(root, 100)
	traverse(root)

	if lookupNode(root, 100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	if lookupNode(root, -100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}
}
