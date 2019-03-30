package main

import (
	"fmt"
	"time"
)

func test_routine() {
	fmt.Println("this is routine")
}

func main()  {
	go test_routine()
	time.Sleep(1 * time.Millisecond)
}

