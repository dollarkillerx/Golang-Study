package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func chandleError(err error,when string) {
	if err != nil {
		fmt.Println(when,err)
		os.Exit(1)
	}
}

func main() {
	conn,e := net.Dial("tcp","127.0.0.1:8580")
	chandleError(e,"net.Dial")
	reader := bufio.NewReader(os.Stdin)// os.Stdin标准输入的读取器
	buffer := make([]byte,1024)

	for {
		lineBytes,_,_ := reader.ReadLine()
		conn.Write(lineBytes)

		n,_:= conn.Read(buffer)
		serverMsg := string(buffer[:n])
		fmt.Println("server:",serverMsg)
	}

	fmt.Println("game over!")
}