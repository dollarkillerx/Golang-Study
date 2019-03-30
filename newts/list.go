package main

import "fmt"

func main()  {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	var grid [4][5]int
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(grid)

	arr3 :=[...]int{2,4,5,9,8}
	for i:=0;i<len(arr3);i++{
		fmt.Println(arr3[i])
	}

	for k,v := range arr3 {
		fmt.Println(k,v)
	}
}