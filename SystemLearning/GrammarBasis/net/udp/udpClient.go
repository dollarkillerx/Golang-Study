package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func chandleError(e error, s string) {
	if e!=nil {
		fmt.Println(e.Error(),s)
		os.Exit(1)
	}
}

func main() {
	conn, e := net.Dial("udp", "127.0.0.1:8051")
	chandleError(e,"net.Dial")
	defer func() {
		conn.Close()
		fmt.Println("呜呼,狡兔死，走狗烹；飞鸟尽，良弓藏；敌国破，谋臣亡..")
	}()

	bytes := make([]byte, 1024)

	for  {
		reader := bufio.NewReader(os.Stdin)
		line, _, _:= reader.ReadLine()

		conn.Write(line)
		n, e := conn.Read(bytes)
		chandleError(e,"conn.Read")
		content := string(bytes[:n])
		fmt.Println("内容:",content)
		if content == "byt" {
			break
		}
	}
}