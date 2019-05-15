package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	//getD()
	postD()
}

func postD() {
	url := "https://httpbin.org/post"
	resp, err := http.Post(
		url,
		"application/x-www-form-urlencoded",
		strings.NewReader("rmn=0.5&name=jeck"),
		)
	handleError(err,"http.Post")
	defer func() {
		resp.Body.Close()
	}()

	bytes,err := ioutil.ReadAll(resp.Body)
	handleError(err,"ioutil.ReadAll")

	fmt.Println(string(bytes))


}

func getD()  {
	resp, err := http.Get("http://www.baidu.com/")
	handleError(err,"http.Get")
	//相应体IO资源,用完就关
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	handleError(err,"ioutil.ReadAll")
	fmt.Println("网页内容",string(bytes))
}

func handleError(e error, s string) {
	if e!=nil {
		fmt.Println(e.Error(),s)
		os.Exit(1)
	}
}

