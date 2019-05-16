package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

func main() {
	//配置并获得一个连接对象的指针
	pool := &redis.Pool{
		//最大闲置链接数
		MaxIdle:20,
		//最大活动链接数,0=无线
		MaxActive:0,
		//闲置链接超时时间
		IdleTimeout:time.Second*100,
		//定义拨号获得链接函数
		Dial: func() (conn redis.Conn, e error) {
			dial, e := redis.Dial("tcp", "172.17.0.2:6379")
			return dial,e
		},
	}

	defer func() {pool.Close()}()

	for i:=0;i<10 ;i++  {
		go getConnFromPlooAndHappy(pool,i)
	}

	//保持主协程存货
	time.Sleep(3*time.Second)
}

func getConnFromPlooAndHappy(pool *redis.Pool,i int)  {
	//通过连接池获得链接
	conn := pool.Get()
	defer conn.Close()
	reply, err := conn.Do("set", "conn"+strconv.Itoa(i), i)
	s,_ := redis.String(reply, err)
	fmt.Println(s)
}