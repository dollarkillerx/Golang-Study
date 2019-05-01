package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n",s,len(s),cap(s))
}

func main() {
	fmt.Println("create slice")
	var s []int // Zero value for slice is nil
	for i := 0; i < 100; i++ {
		s = append(s,2 * i + 1)
		//printSlice(s)
	}
	fmt.Println(s)
	fmt.Println()
	s1 := []int{2,4,6,8}
	printSlice(s1)

	s2 := make([]int,16,25)
	printSlice(s2)

	fmt.Println("Copying slice")
	copy(s2,s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3],s2[4:]...)

	printSlice(s2)

	fmt.Println("Popping from front")
	s2 = s2[1:]
	printSlice(s2)
	fmt.Println("Popping from back")
	s2 = s2[:(len(s2)-1)]
	printSlice(s2)
}