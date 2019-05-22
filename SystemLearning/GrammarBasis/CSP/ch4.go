package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

//生产完 一次性消费

func main() {
	chanShop := make(chan string, 10)
	chanTelphone := make(chan int)
	
	go Producer(chanShop,chanTelphone)//生产
	go Consumer(chanShop,chanTelphone)//消费
	for {
		time.Sleep(time.Second)
	}
}

func Consumer(chanstrings chan string, chanTelphone chan int) {
	//限制生产能力 双核
	runtime.GOMAXPROCS(2)

	for i:=0;i<10 ;i++  {
		itoa := strconv.Itoa(time.Now().Nanosecond())
		chanstrings <- "产品" + itoa

		fmt.Println("生产了产品",itoa)
		//time.Sleep(time.Second)
	}
	close(chanstrings)
	//打电话
	chanTelphone <- 1555555
	close(chanTelphone)
	fmt.Println(1555555,"呼出电话")
}



func Producer(chanstrings chan string, chanTelphone chan int) {
	//限制消费能力,单核消费
	runtime.GOMAXPROCS(1)
	
	//阻塞等电话
	i := <-chanTelphone
	fmt.Println("收到来电",i)
	for c := range chanstrings {
		fmt.Println("消费",c)
	}
	fmt.Println("消费完毕")
}
