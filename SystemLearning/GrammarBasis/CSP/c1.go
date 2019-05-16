package main

import (
	"fmt"
	"time"
)

func main() {
	//test1()
	test3()
}

// 并发
func test1()  {
	//在一个独立的协程中执行匿名函数
	go func() {
		for {
			fmt.Println("im coroutine")
			time.Sleep(time.Second)
		}
	}()

	for  {
		fmt.Println("我的主协程")
		time.Sleep(time.Second)
	}
}

func test3()  {
	//见识一下百万级并发
	for i:=0;i<100;i++ {
		fmt.Println(i)
		go func() {
			fmt.Println("百万级并发")
			time.Sleep(time.Second)
		}( )
	}

	//for {
	//	fmt.Println("我是主协程")
	//	time.Sleep(time.Second)
	//}
}
