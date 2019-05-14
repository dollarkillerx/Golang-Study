package main

import (
	"errors"
	"fmt"
)

func test(a,b int) (int,error) {
	if a>0&&a>0{
		return a+b,nil
	}
	return 0,errors.New("test")
}

func main() {
	var i int
	var e error
	if i, e = test(1, 2);e!=nil{

	}
	fmt.Println(i)
}