package main

import "fmt"

type person struct {
	name string
	age int
}

func (ps person)print()  {
	fmt.Printf("%v 12 12",ps)
}

type bs struct {
	person
	pser int
}

func main() {
	person := person{"dollarkiller",18}
	fmt.Printf("%v",person)
	person.print()
	fmt.Println()
	ps := bs{person,16}
	fmt.Println(ps)
}
