package main

// #cgo CFLAGS: -I${SRCDIR}/exercise-1
// #cgo LDFLAGS: ${SRCDIR}/exercise-1.a
// #include <stdio.h>
// #include <stdlib.h>
// #include <exercise-1.h>
// void printByC() {
// 	printf("Hello World from C!\n");
// }
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	C.printByC()
	additional := C.CString("Hey!")
	defer C.free(unsafe.Pointer(additional))
	C.helloFromC(additional)
	fmt.Println("Hello World from Go!")
}
