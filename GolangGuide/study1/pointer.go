package main

import "fmt"

func test() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Print(a)
}

func main() {
	test()
}
