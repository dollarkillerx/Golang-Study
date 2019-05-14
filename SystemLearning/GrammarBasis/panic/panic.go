package main

import (
	"errors"
)

const (
	ex1 = "异常"
)

func test1(a,b int) int {
	if a>0&&b>0 {
		return a + b
	}else{
		panic(errors.New(ex1))
	}
}

func test2(a,b int) (int,error) {
	if a>0&&b>0 {
		return a + b,nil
	}else{
		return 0,errors.New("不能为负数")
	}
}

func main()  {
	// 在函数结束之前处理panic
	//defer func() {
	//	if err := recover();err!=nil{
	//		switch {
	//			case err == ex1:
	//				fmt.Println("this is 异常")
	//			default:
	//				fmt.Println("default异常处理: ",err)
	//		}
	//	}
	//}()
	//fmt.Println(test1(-4,7))

	if _, e := test2(16, -7);e!=nil{
		println(e.Error())
	}
}
