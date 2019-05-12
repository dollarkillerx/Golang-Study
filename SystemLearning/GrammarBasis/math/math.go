package main

import (
	"fmt"
	"math"
)

//基本数学运算
func main() {
	// 四舍五入
	fmt.Println(math.Round(3.4))
	fmt.Println(math.Round(3.5))
	// 绝对值
	fmt.Println(math.Abs(-335))
	// 乘方
	fmt.Println(math.Pow(2,3))
	// 开方
	fmt.Println(math.Sqrt(9))
}