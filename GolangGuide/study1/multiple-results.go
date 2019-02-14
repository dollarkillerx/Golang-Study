package main

import "fmt"

func swap(s,y string) (string,string) {
	return s,y
}
//函数可以返回任意数量的返回值
func main() {
	a,b :=swap("hello","golang")
	fmt.Println(a,b)
}
