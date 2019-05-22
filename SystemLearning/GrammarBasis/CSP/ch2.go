package main

import (
	"fmt"
	"time"
)

func wrt(ch chan interface{}) {

	for i:=0;i<5;i++{
		ch <- i * i
		fmt.Println("写入",i*i)
		time.Sleep(time.Second)
	}
	//关闭管道,会通知读取协程不再遍历
	close(ch)
	fmt.Println("管道已经关闭")
}

func red(ch chan  interface{}) {
	for x:= range ch{
		fmt.Println("读出:",x)
	}
	fmt.Println("读取协程结束")
}

func main() {
	ints := make(chan interface{},10)
	go wrt(ints)
	go red(ints)
	time.Sleep(6*time.Second)
	fmt.Println("Game Over")
}