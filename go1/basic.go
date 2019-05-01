package main

import (
	"fmt"
)

func enums() {
	//const (
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	golang = 3
	//)
	const (
		cpp = 1 << iota
		java
		python
		_
		golang
	)
	fmt.Println(cpp,java,python,golang)
}

func main() {
	hello := "hello"
	fmt.Print(hello)
	enums()
}