package main

import (
	"fmt"
	"strconv"
	"time"
)

func Add(x,y int)  {
	z := x+y
	fmt.Println(z)
}

func Read(ch chan int) {
	value :=<- ch
	fmt.Println("value:",strconv.Itoa(value))
}

func Write(ch chan int)  {
	ch <- 10
}

func main()  {
	//var ch chan int
	//ch = make(chan int)
	ch :=make(chan int)
	//for i:=1;i<10;i++ {
	//	go Add(i,i)
	//}
	go Read(ch) // 没有读到就阻塞 直到读到为止
	go Write(ch)
	time.Sleep(10*time.Millisecond)
	fmt.Print("end of code")
}

