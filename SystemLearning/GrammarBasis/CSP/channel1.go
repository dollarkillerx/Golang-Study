package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan interface{},3)

	go func() {
		//读
		for {
			if i,ok := <-ch;ok{
				fmt.Println(i)
			}else{
				fmt.Println("数据End")
				return
			}
		}
		//for i :=range ch  {
		//	fmt.Println(i)
		//}
	}()


	mc := make(map[int]string)
	mc[6] = "Test"
	str := []string{"asdasd", "sadasd", "asdasd"}
	//写
	ch<-123
	ch<-mc
	ch<-str
	ch<-"ssss"
	close(ch)

	time.Sleep(time.Second*1)
}
