package main

import "fmt"

const size = 15

type node struct {
	Value int
	Next  *node
}

type hashTable struct {
	Table map[int]*node
	Size  int
}

func hashFunction(i, size int) int {
	return (i % size)
}

func insert(hash *hashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *hashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func lookup(hash *hashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	if t := hash.Table[index]; t != nil {
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main() {
	table := make(map[int]*node, size)
	hash := &hashTable{Table: table, Size: size}
	fmt.Println("Number of spaces:", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	for i := 115; i < 125; i++ {
		if lookup(hash, i) {
			fmt.Printf("%d is in the hash table!\n", i)
		} else {
			fmt.Printf("%d is not in the hash table!\n", i)
		}
	}
}
