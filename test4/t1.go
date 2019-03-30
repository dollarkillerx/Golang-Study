package main

import "fmt"

func main()  {
	primes := [6]int{2,6,5,8,8,4}
	var mmf []int = primes[2:5]
	fmt.Println(mmf)
}
