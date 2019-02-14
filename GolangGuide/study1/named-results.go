package main

import "fmt"

func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum -x
	return 
}
//命名返回值  golang返回值可以命名 没有参数的return 返回结果当前值`


func main() {
	fmt.Println(split(17))
}
