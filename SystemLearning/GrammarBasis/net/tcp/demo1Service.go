package main

import (
	"fmt"
	"net"
	"os"
)

/**
服务端在本级的8888端口建立tcp监听
为接入的每一个客户端,开辟一条独立的协程
循环接收客户端消息,不管客户端说什么,都自动恢复"已阅读xxx"
如果客户端"im off" 回复bye 并断开链接
 */

// 错误处理
func handlingErrors(err error,sign string)  {
	if err!=nil {
		fmt.Println(err.Error(),sign)
		os.Exit(1)//非0代表异常
	}
}

func main() {
	listener, e := net.Listen("tcp", "127.0.0.1:8085")
	handlingErrors(e,"Listen tcp")


	//接入每一个客户端,开辟协程
	for {
		conn, e := listener.Accept()
		handlingErrors(e,"listener.Accept")
		go chatWith(conn)
	}
}

func chatWith(conn net.Conn) {
	buffer := make([]byte,1024)
	for {
		n, err := conn.Read(buffer)
		handlingErrors(err,"conn.Read(buffer)")

		clientMsg := string(buffer[0:n])
		fmt.Printf("收到%v的消息:%s\n",conn.RemoteAddr(),clientMsg)
		if clientMsg == "im off" {
			conn.Write([]byte("byt"))
			break
		}

		conn.Write([]byte("已阅读" + clientMsg))
	}
	defer func() {
		conn.Close()
		fmt.Println("client disconnected!")
	}()
}