package main

import (
	"fmt"
	"net"
	"os"
)

func handleError(err error,when string) {
	if err != nil {
		fmt.Println(when,err)
		os.Exit(1)
	}
}

func ioWithConn(conn net.Conn) {
	buffer := make([]byte,1024)
	for {
		n,err := conn.Read(buffer)
		handleError(err,"conn.Read(buffer)")
		clientMsg := string(buffer[:n])
		fmt.Printf("received cleint:%v msg:%s\n",conn.RemoteAddr(),clientMsg)
		if clientMsg=="im off" {
			conn.Write([]byte("BYE!"))
			break
		}
		conn.Write([]byte("msg received" + clientMsg))
	}
	fmt.Println("client disconnected!")
}

func main() {
	listener, e := net.Listen("tcp", "127.0.0.1:8580")
	handleError(e,"net.Listen")
	fmt.Println("listening run 127.0.0.1:8580 ...")

	for {
		conn,e := listener.Accept()
		handleError(e,"listener.Accept")
		go ioWithConn(conn)
	}
}