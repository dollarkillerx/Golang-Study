package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func apply(op func(int,int) int,a,b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args" + "(%d,%d)",opName,a,b)
	return op(a,b)
}

func add(a,b int) int {
	return a + b
}

func main() {
	fmt.Println(apply(add,14,20))
}