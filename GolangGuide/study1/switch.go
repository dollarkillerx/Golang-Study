package main

import (
	"fmt"
)

func eval(a,b int,op string) int {
	var result int
	switch op {
	case "+":
		result = a+b
	case "-":
		result = a-b
	case "*":
		result = a*b
	case "/":
		result = a/b
	default:
		panic("unsupported operator" + op)
	}
	return result
}

func grade(score int) string {
	g := ""
	switch {
	case score<0 || score>100:
		panic(fmt.Sprintf("Wrong score:%d",score))
	case score>80:
		g = "a"
	case score>60:
		g = "b"
	case score<60:
		g = "f"
	}
	return g
}

func main()  {
	a := eval(1,2,"-")
	fmt.Print(a)
}
