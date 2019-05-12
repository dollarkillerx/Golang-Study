package enc

import (
	"fmt"
	"testing"
)

type Emp struct {
	name string
	age int
}

func (e *Emp) pri() {
	fmt.Println(e.name,e.age)
}

func TestStruct(t *testing.T){
	em := Emp{"dolalr",16}
	em.pri()
}