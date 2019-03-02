package main

import "fmt"

func add(a *[5]int) {
	a[1] = 123
}

func main() {
	var arr1 [5]int // 初始值0
	arr2 := [3]int{1,3,5} // :=必先初始化
	arr3 := [...]int{2,4,6,8,10} // 定义切片
	var grid [4][5]int // 定义四行五列

	fmt.Println(arr1,arr2,arr3,grid)
	for i := 0; i<len(arr3);i++  {
		fmt.Print(arr3[i]," ")
	}

	for k,v := range arr3 { // 下标 value  if 只有k就只要下标
		fmt.Println("key: ",k,"value: ",v)
	}

	for _,v := range arr3 {
		fmt.Print(v)
	}
	add(&arr3)
	fmt.Println(arr3)
}
