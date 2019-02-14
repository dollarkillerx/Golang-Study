package main

import "fmt"

func add(x,y int) int {
	return (x+y)
}
//当两个连续的函数命名参数是统一类型,则除了最后一个类型外其他都快省略`

func main(){
	fmt.Print(add(16,26))
}
