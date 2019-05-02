package main

import "fmt"

func div(a, b int) (int, int) {
	return a / b, a % b
}

func add(a, b int) int {
	return a + b
}

func apply(add func(int, int) int, a, b int) int {
	return add(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(div(15, 7))
	fmt.Println(fmt.Errorf("ERROR%s", "sdsd"))
	fmt.Println(apply(add, 1, 5))
	fmt.Println(sum(1, 2, 45, 4, 5, 4, 87, 8, 456, 45, 6, 167, 1, 56, 74, 87))
}
