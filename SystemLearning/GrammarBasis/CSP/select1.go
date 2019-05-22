package main

import "fmt"

func main() {
	ch0 := make(chan int,0)
	ch1 := make(chan int,1)
	ch2 := make(chan int,2)

	ch2 <- 123
	ch2 <- 456


	select { //随机选一个能跑的
		case ch0 <- 123:
			fmt.Println("ch0")
		case i := <-ch1:
			fmt.Println("ch1",i)
		case ch2 <- 123:
			fmt.Println("ch2")
		default:
			fmt.Println("End")
	}
}