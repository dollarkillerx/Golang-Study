package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, e := redis.Dial("tcp", "172.17.0.2:6379")
	handleError(e)
	defer func() {
		conn.Close()
	}()

	//执行redis命令,get return data
	reply, e := conn.Do("Get", "name")
	handleError(e)
	fmt.Printf("type=%T,value=%v\n",reply,reply)

	//类型装换
	s, e := redis.String(reply, e)
	handleError(e)
	fmt.Println(s)
}

func handleError(e error) {
	if e != nil {
		panic(e.Error())
	}
}
