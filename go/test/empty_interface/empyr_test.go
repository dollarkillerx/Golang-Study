package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}){
	//if i,ok:=p.(int);ok{
	//	fmt.Println("int",i)
	//}
	//if i,ok:=p.(string);ok{
	//	fmt.Println("string",i)
	//}
	switch v:=p.(type) {
	case int:
		fmt.Println("this is int",v)
	case string:
		fmt.Println("this is string",v)
	}
}

func TestEmptyInterfaceAssertion(t *testing.T)  {
	DoSomething("1616")
}