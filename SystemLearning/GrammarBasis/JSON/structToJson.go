package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
	Rmb float64
	Sex bool
	Hobby []string
}

func structToJson()  {
	/**
	使用Json包
	*/
	qianGe := Person{"于谦", 50, 123.45, true, []string{"抽烟", "喝酒", "烫头"}}
	byt,err := json.Marshal(qianGe)//吧任意数据,转为json
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(string(byt))
}

func main() {
	structToJson()
}