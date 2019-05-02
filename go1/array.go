package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 9}
	arr3 := [...]int{2, 4, 5, 6, 90}
	fmt.Println(arr1, arr2, arr3)
	printArray(&arr3)
	fmt.Println()
	for _, v := range arr3 {
		fmt.Println(v)
	}
}
