package main

import "fmt"
var x,y = 1,2
func main() {
	var i,j int = 1, 2
	c,python,java,php := true,true,false,"ok!"
	//在函数中可以使用 `:=` 简洁赋值语句在明确类型的地方可以用于替代var定义  BUT在函数体外必须以关键字开始 `:=`不能用于函数体外
	fmt.Println(x,y,i,j,c,python,java,php)
}
