package main

import "fmt"

func main()  {
	arr :=[...]int{0,1,2,3,4,5,6,7}
	s := arr[2:6]
	fmt.Println(s)
}