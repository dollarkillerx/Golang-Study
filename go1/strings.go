package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "我爱中国"
	fmt.Println(s,len(s))
	for _,b := range []byte(s) {
		fmt.Printf("%X ",b)
	}
	fmt.Println()
	for i,ch := range s { //ch is a rune
		fmt.Printf("(%d,%X) ",i,ch)
	}
	fmt.Println()

	s = "123, 145 ,64"
	fmt.Println(strings.Count(s,"1"))
	fmt.Println(strings.Fields(s))
}
