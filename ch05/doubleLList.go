package main

import (
	"errors"
	"fmt"
)

var errNodeAlreadyExist = errors.New("Node already exist in list")
var errNodeDoesNotExist = errors.New("Node does not exists")

type node struct {
	Value int
	Prev  *node
	Next  *node
}

const minInt = -int(^uint(0)>>1) - 1

// Returns new HEAD pointer after inserting element
func addNode(t *node, v int) (*node, error) {
	if t == nil {
		return &node{v, nil, nil}, nil
	}

	temp := &node{minInt, nil, t}
	t = temp
	for t.Next != nil {
		if t.Next.Value == v {
			return temp.Next, errNodeAlreadyExist
		} else if t.Next.Value > v {
			var inserted *node
			if t == temp {
				inserted = &node{v, nil, t.Next}
			} else {
				inserted = &node{v, t, t.Next}
			}
			t.Next.Prev = inserted
			t.Next = inserted
			break
		}
		t = t.Next
	}

	if t.Next == nil {
		t.Next = &node{v, t, nil}
	}

	return temp.Next, nil
}

func deleteNode(t *node, v int) (*node, error) {
	if t == nil {
		return nil, errNodeDoesNotExist
	}

	head := t
	for t != nil {
		if t.Value == v {
			if t == head {
				t.Next.Prev = nil
				return t.Next, nil
			}
			t.Next.Prev = t.Prev
			t.Prev.Next = t.Next
			return head, nil
		}
		t = t.Next
	}

	return head, errNodeDoesNotExist
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

func reverse(t *node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t.Next != nil {
		t = t.Next
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Prev
	}
	fmt.Println()
}

func size(t *node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupNode(t *node, v int) bool {
	if t == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func main() {
	var root *node
	var err error
	fmt.Println(root)
	traverse(root)
	if root, err = addNode(root, 1); err != nil {
		fmt.Println(err)
	}
	if root, err = addNode(root, 1); err != nil {
		fmt.Println(err)
	}
	traverse(root)
	root, _ = addNode(root, 10)
	root, _ = addNode(root, 5)
	root, _ = addNode(root, 0)
	root, _ = addNode(root, 0)
	root, _ = addNode(root, 100)
	root, _ = addNode(root, -3)
	traverse(root)
	root, _ = deleteNode(root, 0)
	root, _ = deleteNode(root, 0)
	fmt.Println("Size:", size(root))
	traverse(root)
	reverse(root)
}
