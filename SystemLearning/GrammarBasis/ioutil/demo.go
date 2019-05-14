package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/TestFile1.test"
	var content string
	if bytes,err := ioutil.ReadFile(file);err!=nil{
		panic(err.Error())
	}else{
		content = string(bytes)
	}

	var numCount int
	for v := range content {
		if v >= '0'&& v <= '9' {
			numCount ++
		}
	}
	fmt.Print("numCount",numCount)

}