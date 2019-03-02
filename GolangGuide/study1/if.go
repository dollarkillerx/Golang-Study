package main 

import "fmt"

func bounded(v int) int {
	if v > 100{
		return 100
	} else if v<0 {
		return 0
	}else {
		return v
	}
}

func main() {
	fmt.Print(bounded(16))
}
