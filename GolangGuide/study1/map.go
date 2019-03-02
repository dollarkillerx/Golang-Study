package main

import "fmt"

func main()  {
	m := map[string]string {
		"name": "dollarkiller",
		"age": "18",
		"course": "golang",
	}
	m1 := map[string]int {
		"age": 16,
	}
	fmt.Println(m," ",m1)
	cases,ok := m["case"]
	fmt.Println(cases,ok)
}