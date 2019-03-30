package main

import "fmt"

func fo()  {
	sum := 0
	for i:=1;i<=100;i++ {
		fmt.Println(sum)
		sum += i
	}
	fmt.Printf("蛤%s",sum)
}

func main()  {
	fo()
}