package main

import "fmt"

func eval(a,b int,op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	default:
		panic("un")
	}
}

func main() {
	a := eval(15,46,"+")
	fmt.Println(a)
}
