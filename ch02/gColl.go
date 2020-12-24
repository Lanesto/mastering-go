package main

import (
	"fmt"
	"runtime"
	"time"
)

var latestMemStats runtime.MemStats

func printMemStats() {
	runtime.ReadMemStats(&latestMemStats)
	ms := latestMemStats
	fmt.Printf("showing memory statistics on %s\n", time.Now().Format(time.StampMilli))
	fmt.Println("mem.Alloc:", ms.Alloc)
	fmt.Println("mem.TotalAlloc:", ms.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", ms.HeapAlloc)
	fmt.Println("mem.NumGC:", ms.NumGC)
	fmt.Println("-------------------------------")
}

func main() {
	printMemStats()

	for i := 0; i < 10; i++ {
		s := make([]byte, 2000000)
		if s == nil {
			fmt.Println("error: memory allocation failed")
		}
	}
	printMemStats()

	for i := 0; i < 10; i++ {
		s := make([]byte, 10000000)
		if s == nil {
			fmt.Println("error: memory allocation failed")
		}
		time.Sleep(1 * time.Second)
	}
	printMemStats()
}
