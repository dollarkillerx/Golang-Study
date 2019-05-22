package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//商店管道:生产者和消费者传递产品
	chanStorage := make(chan string, 100)//仓库
	chanShop := make(chan string,100) //商店

	//生产者和消费者
	go producer(chanStorage) //生产
	go wuliu(chanStorage,chanShop) //物流
	go consumer(chanShop)	//消费

	for  {
		
	}
}

//生产者
func producer(chanstrings chan string) {
	for i:=0;i<10 ;i++  {
		itoa := strconv.Itoa(time.Now().Nanosecond())
		chanstrings <- "产品" + itoa
		fmt.Println("生产了产品:",itoa)
		time.Sleep(time.Second)
	}
	close(chanstrings)
	fmt.Println("生产全部完毕")
}
//物流
func wuliu(chanstrings,chanshop chan string) {
	for p:= range chanstrings{
		fmt.Println("物流完成转运",p)
		chanshop <- p
	}
	close(chanshop)
	fmt.Println("物流全部完毕")
}
// 消费
func consumer(chanShop chan string) {
	for i:= range chanShop {
		fmt.Println("消费了",i)
	}
	fmt.Println("消费全部完毕")
}









