package main

import (
	"fmt"
	"os"
)

func main() {
	// 获得当前工作路径
	dir,error := os.Getwd()
	if error!=nil {
		panic(error.Error())
	}
	fmt.Println(dir)
	// 获得指定的环境变量
	paths := os.Getenv("Path")
	goroot := os.Getenv("GOROOT")
	fmt.Println(paths)
	fmt.Println(goroot)
	hostname,err := os.Hostname()
	if err!=nil {
		panic(err.Error())
	}
	fmt.Println(hostname)

	fmt.Println()

	// 获得文件的状态
	fileinfo,err := os.Stat("/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/goto/goto.go")
	if err!=nil {
		panic(err.Error())
	}
	// 是否是文件夹
	fmt.Println(fileinfo.IsDir())
	// 权限
	fmt.Println(fileinfo.Mode())
	// 修改时间
	fmt.Println(fileinfo.ModTime())
	// 文件名称
	fmt.Println(fileinfo.Name())
	// 文件大小
	fmt.Println(fileinfo.Size())

}

