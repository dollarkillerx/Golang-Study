package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	answer := myRand.Intn(1000)

	for {
		// 接收用户输入
		fmt.Println("请输入你的猜想:")
		var guess string
		fmt.Scan(&guess)

		// 转换为整数
		guessNum,_ := strconv.Atoi(guess)

		// 返回比较
		switch {
		case guessNum>answer:
			fmt.Println("猜大了")
		case guessNum<answer:
			fmt.Println("猜小了")
		default:
			fmt.Println("猜到了")
			break
		}
	}
}
