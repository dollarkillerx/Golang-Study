# 回顾Golang复习
```
func main() {
	if len(os.Args)>1 {
		fmt.Println("hello",os.Args[1]) // 返回命令行参数
	}
	os.Exit(-1) //返回程序退出时的状态
}
```
#### 编写测试程序
- 源码文件以_test结尾: xxx_test.go
- 测试方法名以Test开头: `func TestXXX(t *testing.T){xxx}`