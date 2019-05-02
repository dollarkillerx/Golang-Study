package main

import "fmt"

func adder() func(int) int {
	sum := 0 // 自由变量  会被保存下来
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	a := adder()
	for i:=0;i<10;i++  {
		fmt.Println(a(i))
	}
}
