package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [...]int{0, 1, -2, 3, 4}
	memSize := unsafe.Sizeof(array[0])

	var head *int
	nextAddr := uintptr(unsafe.Pointer(&array[0]))
	for i := 0; i < len(array); i++ {
		head = (*int)(unsafe.Pointer(nextAddr))
		fmt.Print(*head, " ")
		nextAddr = uintptr(nextAddr + memSize)
	}
	// Following will pop trash value; it is not element of array
	head = (*int)(unsafe.Pointer(nextAddr))
	fmt.Print(*head)

	fmt.Println()
}
