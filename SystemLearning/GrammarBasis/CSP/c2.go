package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	test2()
}

func test2()  {
	var a [10]int
	for i:=0; i<10;i++  {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched()//交出控制器
				//fmt.Printf("hello from goroutine %d \n",i)
			}
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(a)

}

func test11()  {
	go func() {
		for {
			fmt.Println("我是子携程")
			time.Sleep(time.Second)
		}
	}()

	for {
		fmt.Println("我是主携程................................")
		time.Sleep(time.Second)
	}
}