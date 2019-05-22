package main

import (
	"fmt"
	"time"
)

func chanDemo()  {
	// var c chan int // c == nil
	c := make(chan int)
	go worker(c)
	c <- 1  //发数据
	c <- 2

	time.Sleep(time.Millisecond)
}

func worker(c chan int) {
	for {
		n := <- c //收数据
		fmt.Println(n)
	}
}

func test1()  {
	fmt.Println("ttt")
}

func main() {
	//chanDemo()

	for i:=0;i<=100;i++ {
		go test1()
	}

	for  {
		fmt.Println("主协程")
		time.Sleep(time.Second)
	}
}