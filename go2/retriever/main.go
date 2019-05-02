package main

import (
	"fmt"
	"go2/retriever/mock"
	"go2/retriever/queue"
	"go2/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func post(poster Poster) {
	poster.Post("https://www.baidu.com",
		map[string]string{
			"name":"dollarKiller",
			"course":"golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster)  {
	s.Get("https://www.baidu.com")
	s.Post("https://www.baidu.com",
		map[string]string{
			"name":"dollarKiller",
			"course":"golang",
		})
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake dollarkiller"}
	fmt.Println(download(r))
	fmt.Println("123132")
	r = real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	fmt.Println(download(r))
	fmt.Println()
	fmt.Printf("%T {:} %v\n",r,r)
	fmt.Println()

	q := queue.Queue{1}
	q.Push(2)
	q.Push("你好")
	q.Push(false)
	fmt.Println(q)
}