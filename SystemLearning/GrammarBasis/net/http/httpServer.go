package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	html := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/net/http/index.html"
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		bytes, _ := ioutil.ReadFile(html)
		writer.Write(bytes)
	})

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("请求的方法: " + request.Method))
		writer.Write([]byte("内容长度: " + strconv.Itoa(int(request.ContentLength))))
		writer.Write([]byte("主机地址: " + request.Host))
		writer.Write([]byte("请求的协议: " + request.Proto))
		writer.Write([]byte("请求的远程地址: " + request.RemoteAddr))
		writer.Write([]byte("请求的路由: " + request.RequestURI))
		fmt.Println(request.Header["User-Agent"])
	})

	http.ListenAndServe("127.0.0.1:8580",nil)
}
