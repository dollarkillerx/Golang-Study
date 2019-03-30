package main

import "fmt"

func huan(a,b *int)  {
	*a,*b = *b,*a
}

func main() {
	a,b := 5,6
	huan(&a,&b)
	fmt.Printf("a:%d,b:%d",a,b)
}