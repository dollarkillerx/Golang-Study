package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)>1 {
		fmt.Println("hello",os.Args[1]) // 返回命令行参数
	}
	os.Exit(-1) //返回程序退出时的状态
}
