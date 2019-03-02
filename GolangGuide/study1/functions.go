package main

import "fmt"

func add(x int,y int) int {
	return x + y
}
//函数可以没有参数或者接受多个参数

func test(number ...int) int {
	s := 0
	for i:=range number {
		s += number[i]
	}
	return s
}

func applay(op func(int,int) int,a,b int) int {
	return op(a,b)
}

func main() {
	//fmt.Print(add(16,26))
	//fmt.Print(test(1,2))
	//fmt.Print(test(12,5445,89,787))
	fmt.Print(applay(add,16,15))
	fmt.Print(applay(func(i2 int, i22 int) int {
		return i2 * i22
	},15,16))
}
