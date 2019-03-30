package main

import "fmt"

func zz()  {
	var a = 2
	var pa *int = &a

	*pa = 23
	fmt.Println(a)
}

func main()  {
	zz()
}

