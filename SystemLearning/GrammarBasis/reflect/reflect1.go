package main

import (
	"fmt"
	"reflect"
)

func main() {
	testValue()
}

type Animal struct {
	lift int
}

type People struct {
	name string
	age int
	animal Animal
}

func (People) hello()  {
	fmt.Println("Hello")
}

func testValue()  {
	p := People{"张傻大",20,Animal{123}}
	//p.hello()
	//typeApi(p)
	valueApi(p)
}

func typeApi(obj interface{})  {
	oType := reflect.TypeOf(obj)
	fmt.Println(oType) //main.People

	// 原始类型
	oKind := oType.Kind()
	fmt.Println(oKind) //struct

	//类型名称
	fmt.Println(oType.Name())

	//属性和方法 个数
	fmt.Println(oType.NumField())//属性
	fmt.Println(oType.NumMethod())//方法

	fmt.Println("全部属性:")
	for i:=0;i<oType.NumField();i++{
		structField := oType.Field(i)
		fmt.Println(structField.Name,structField.Type)
	}

	fmt.Println("全部方法")
	for i:=0;i<oType.NumMethod();i++ {
		method := oType.Method(i)
		fmt.Println(method.Name,method.Type)
	}
}

func valueApi(o interface{})  {
	oValue := reflect.ValueOf(o)
	fmt.Println(oValue)

	fmt.Println("获得oValue的原始类型")
	fmt.Println(oValue.Kind()) //struct

	fmt.Println("查看p所有属性的值")
	for i:=0;i<oValue.NumField();i++ {
		fValue := oValue.Field(i)
		fmt.Println(fValue.Interface())// 属性值的正摄
	}
}