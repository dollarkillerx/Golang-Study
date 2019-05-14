package main

import (
	"fmt"
	"time"
)

func main()  {
	// 设置时区
	l,_ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().In(l))

	//str转时间
	formatTimeStr := "2019-05-12"
	formatTime,err := time.Parse("2006-01-02",formatTimeStr)
	if err==nil{
		fmt.Println(formatTime.Unix()) //打印结果：2017-04-11 13:33:37 +0000 UTC
	}

	nowTime := time.Now().In(l)
	fmt.Println(nowTime.Date())
	fmt.Println(nowTime.Unix())// 获取时间戳

	timeLayout := "2006-01-02"  //转化所需模板
	//时间戳转str
	time := time.Unix(nowTime.Unix(),0).Format(timeLayout)
	fmt.Println(time)


}