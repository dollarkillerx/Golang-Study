package main

import (
	"fmt"
	"strings"
)



func main() {
	st := "helLo Golang"
	sf := "golang"
	ss := "xbf"
	zst := "来玩Golang吧"
	zss := '来'

	// 是否包含子串
	fmt.Println(strings.Contains(st,sf))
	// 是否包含任意一个
	fmt.Println(strings.ContainsAny(st,ss))
	// 是否包含字符
	fmt.Println(strings.ContainsRune(zst,zss))

	// 子串在父串出现的序号 没有返回-1
	fmt.Println(strings.Index(st,sf))
	fmt.Println(strings.IndexRune(zst,zss)) // 支持中文

	//大小写转换
	fmt.Println(strings.ToLower(st))
	fmt.Println(strings.ToUpper(st))
	fmt.Println(strings.Title(st))

	test_str := "   hello google  "
	// 裁剪空格
	fmt.Println(strings.TrimSpace(test_str))// 头尾的空格
	fmt.Println(string(test_str[3]))
	test_str = "hello google"
	fmt.Println(strings.TrimPrefix(test_str,"ll"))//去掉前缀 (返回后面部分,如果没有前缀就返回原来的)
	fmt.Println(strings.TrimSuffix(test_str,"le"))//去掉后缀 (返回后面部分,如果没有后缀就返回原来的)
	test_str = "helloBellohelloBhello"
	fmt.Println(strings.Trim(test_str,"hello")) // 去掉头尾包含的
	fmt.Println(strings.TrimLeft(test_str,"hello")) // 去left部分包含的
	fmt.Println(strings.TrimRight(test_str,"hello")) // 去right部分包含的
	// 自定义过滤 TrimFunc(s string ,f func(rune) bool) 传一个函数,返回bool
	fmt.Println(strings.TrimFunc(test_str,filter))

	// string 转 array
	test_str = "hello.mp4,hello.jpg,hbbc.prc,"
	strs := strings.Split(test_str,",")
	strs = strs[:(len(strs)-1)]
	for k,v:=range strs {
		fmt.Println(k,":",v)
	}

	// 拼接
	slice := []string{"hello","golang","gg","ssc"}
	ssr := strings.Join(slice,",")
	fmt.Println(ssr)


}
func filter(char rune) bool {
	if char=='o'||char=='l' {
		return true
	}
	return false
}


