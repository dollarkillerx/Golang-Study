package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 错误处理
func chandlingErrors(err error,sign string)  {
	if err!=nil {
		fmt.Println(err.Error(),sign)
		os.Exit(1)//非0代表异常
	}
}

func main() {
	conn, e := net.Dial("tcp", "127.0.0.1:8085")
	chandlingErrors(e,"net.Dial")
	buffer := make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)

	for {
		lineBytes, _, _ := reader.ReadLine()

		conn.Write(lineBytes)

		n, e := conn.Read(buffer)
		chandlingErrors(e,"conn.Read(buffer)")
		bytes := string(buffer[0:n])
		fmt.Println("服务端:",bytes)

		if  bytes == "byt"{
			break
		}
	}

	defer func() {
		fmt.Println("over")
		conn.Close()
	}()

}