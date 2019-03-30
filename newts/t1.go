package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename= "./main.go"
	//contents,err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	//fmt.Println(contents)
	//	fmt.Printf("%s\n",contents)
	//}
	if contents, err := ioutil.ReadFile(filename); err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}

}