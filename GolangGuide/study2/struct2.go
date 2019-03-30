package main

import "fmt"

type Integer struct {
	age int
}

func (age *Integer) add ( ad int) {
	age.age += ad
	fmt.Println(age)
}

func main() {
	i1 := Integer{12}
	fmt.Println(i1)
	i1.add(15)
	fmt.Println(i1)
}