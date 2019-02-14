package main

import "fmt"

var cpp, python,php,java bool
//定义在func外面为全局变量
func main() {
	var i int
	//定义在func内为局部变量 
	fmt.Println(i,cpp,python,php,java)
}
