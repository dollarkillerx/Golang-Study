package main

import "fmt"

type Animal interface {
	Fly()
	Run() bool
}

type Bird struct {

}

func (bird Bird) Fly()   {
	fmt.Print("Bird")
}

func (bird Bird) Run() bool {
	return true
}


func main() {
	var animal Animal
	bird := new(Bird)
	animal = bird
	animal.Fly()
	fmt.Println(animal.Run())
	var v1 interface{}
	v1 = "sadsadasd"
	fmt.Println(v1)
}