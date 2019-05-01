package main

import "fmt"

func returnTest(a,b int) (int,int) {
	return a,b
}

func main()  {
	m := map[string]string {
		"name":"dollarkiller",
		"course":"golang",
		"site":"imooc",
		"quality":"notbad",
	}
	fmt.Println(m["name"])
	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int // m2 = nil
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for i:=range m{
		fmt.Println(i," : ",m[i])
	}
	fmt.Println("Getting values")
	fmt.Println()
	if 	couresName,ok := m["cours e"];ok{
		fmt.Println(couresName)
	}else{
		fmt.Println("k is empty")
	}

	a,_ := returnTest(1,5)
	fmt.Println(a)
	fmt.Println("Delete values")
	delete(m,"name")
	fmt.Println(m)
}