package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 解析得到udp地址
	addr, e := net.ResolveUDPAddr("udp", "127.0.0.1:8051")
	shandleError(e,"net.ResolveUDPAddr")
	conn, e := net.ListenUDP("udp", addr)
	//在url地址上建立udp监听,获得链接
	shandleError(e,"net.ListenUDP")
	bytes := make([]byte, 1024)
	for  {
		//从链接中读出内容,丢入缓冲区
		i, udpAddr, e := conn.ReadFromUDP(bytes)
		shandleError(e,"conn.ReadFromUDP")
		content := string(bytes[:i])

		fmt.Printf("读到来自%v的内容:%s\n",udpAddr,content)

		if content == "byt" {
			//向客户端写出内容
			conn.WriteToUDP([]byte("byt"),udpAddr)
		}else{
			//向客户端写出内容
			conn.WriteToUDP([]byte("已读" + content),udpAddr)
		}
	}
}

func shandleError(e error, s string) {
	if e!=nil {
		fmt.Println(e.Error(),s)
		os.Exit(1)
	}
}