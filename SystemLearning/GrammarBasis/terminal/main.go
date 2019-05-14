package main

import (
	"flag"
	"fmt"
)

func main() {
	//cc := os.Args
	//for k,v := range cc {
	//	fmt.Println(k,':',v)
	//}

	name := flag.String("name","default","comment")

	//解析获取参数,丢入参数的指针中
	flag.Parse()
	fmt.Print(*name)
}
