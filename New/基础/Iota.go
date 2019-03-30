package main

import "fmt"

const (
	m1 = iota + 15
	//自动沿用排头表达式 iota 自动递增  iota初始值0
	m2
	m3
	m4
	m5
)

func main() {
	fmt.Print(m5)
}
