package main

import "fmt"

// 自定义非法范围异常 异常
type InvalidRadiusError struct {
	s string
}

func (e * InvalidRadiusError) Error() string {
	return e.s
}

func newInvalidRadiusError(text string) error {
	return &InvalidRadiusError{text}
}

func add(a,b int) int  {
	if a>0&&b>0 {
		return a + b
	}else{
		panic(newInvalidRadiusError("超出范围"))
	}
}

func main() {
	defer func() {
		if err := recover();err != nil{
			fmt.Println(err)
		}
	}()

	fmt.Println(add(15,-8))
}


