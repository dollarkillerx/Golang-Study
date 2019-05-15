package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

/**
所有商品都有一些共性:品名,价格
个性无千无万,封装三种商品
睡衣给出一个商品的切片,将每件商品的所有属性输出到JSON文件(品名.json)
用反射实现
 */

type Computer struct {
	Name string
	Price float64
	Cpu string
	Memory int
	Disk int
}

type TShirt struct {
	Name string
	Price float64
	Color string
	Size int
	Sex bool
}

type Car struct {
	Name string
	Price float64
	Cap int
	Power string
}

func main() {
	products := make([]interface{},0)

	products = append(products,Computer{"air",10000,"ryzen3",32,1024})
	products = append(products,TShirt{"ryzen",100,"ryzen",40,false})
	products = append(products,Car{"乔治巴顿",800000,6,"柴油"})
	//test1(&products)
	test2(products)
}

func test1(o *[]interface{})  {
	ref := reflect.TypeOf(o)
	name := ref.Name()
	jsonEncode(o,name)
}

func test2(o []interface{})  {
	for _,i := range o {
		//fmt.Println(i)
		//of := reflect.TypeOf(i)
		//fmt.Println(of.Name())

		of := reflect.ValueOf(i)
		name := of.FieldByName("Name").Interface()
		fmt.Println(name)
	}
}

func jsonEncode(o *[]interface{},fileName string) {
	filePath := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/reflect/test"
	file,err := os.OpenFile(filePath,os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0666)
	if err!=nil {
		panic(err.Error())
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(o)
	if err!=nil {
		panic(err.Error())
	}
	defer func() {
		fmt.Println("编码成功")
		file.Close()
	}()
}