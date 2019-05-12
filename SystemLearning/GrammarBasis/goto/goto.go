package main

import "fmt"

func main() {
	fmt.Println("日照香炉生紫烟")
	fmt.Println("遥看瀑布挂前川")
	goto GAMEOVER
	fmt.Println("飞流直下三千尺")
	LARS:
	fmt.Println("疑是银河落九天")
	return
	GAMEOVER:
	fmt.Println("G ver")
	goto LARS
}