package main

import "fmt"

func convArrToMap(arr []string) map[int]string {
	result := make(map[int]string)
	for idx, val := range arr {
		result[idx] = val
	}
	return result
}

func main() {
	arr := []string{"Hello", "Go", "Programming!"}
	convMap := convArrToMap(arr)
	fmt.Println(arr, "converted to", convMap)
}
