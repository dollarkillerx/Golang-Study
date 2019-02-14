# Golang-Study
Golang Study
### gob
``` 
vim /etc/profile
export GOROOT=/opt/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/goproject
```
### hello golang
``` 
package main

import "fmt"

func main(){
    fmt.Print("Helo Golang!")
}
```
go run hello.go 编译并运行

### $GOPATH 约定三个子目录
- src 存放源代码(.go .c .h .s)
- pkg 编译后生成的文件 (.a)
- bin 编译后生成的可执行文件
``` 
go get github.com/beego/bee  #网站/开发者/package
$GOPATH/src/github.com/beego/bee
```

### Basics
- 特色 select go chan
- 基础结构
``` 
import "main" //程序所属包
import "fmt" //导入依赖包
const MANE string = "dollarkiller" //const常量
var age = 20 //全局变量

//一般类型声明
type ageInt int
//结构体声明
type Dollar struct {
}
//声明接口
type IDollar interface {
}
//函数定义
func Dor() {
    fmt.Print("DollarKiller")
}
//main()入口函数
func main() {
    Dor()
    fmt.Println("GOLANG")
}

```
- package
``` 
    package <pkgName>
    package main
    >>>
    main.main()函数独立程序入口点
    生成go可执行文件,必须要有main的package,且下必须要有main()函数
    同一个路径下只能存在一个package,一个package可以拆分成多个源文件
```
- import
>在 Go 中，首字母大写的名称是被导出的。
``` 
import 只能导入有package的包
不能导入源文件没有用到的包,否则会报错
语法格式1:
import "package1"
import "package2"
语法格式2:
import (
    "package1"
    "package2"
    ...
)
```
    - import 原理
    ```
    1.如果一个main导入其他的包，包将被顺序导入；
    2.如果导入的包中依赖其他包(包B),会首先导入包B,然后初始化B包中的常量变量,最后执行B包中的init()
    3.所有包导入完成后才会对main中的常量和变量进行初始化,然后指向main中的init函数(如何存在),最后执行main函数
    4.如果一个包被导入多次这该包只会被导入一次 
    ```
![import](./readme/golang.png)    
- var 定义 var varName type 蛤,和es5好像  定义了不用也会报错
``` 
func main(){
	var x int
	x = 1
	y,z := "DollarKiller","golang" //无法用于函数外部
	fmt.Printf("%d,%s,%s",x,y,z)
}
```
- const 常量
- int8 int16 int32 int64
- byte
- uint8 uint16 uint32 uint64 unsigned 无符号
- float32 float64
- complex 复数
- string
- array元组(静态) slice列表(动态)
- map hash字典
>任何数据类型以value形式输出fmt.Printf("%v",x)

### golang指南

#### 基础
- 多值返回
``` 
package main

import "fmt"

func swap(s,y string) (string,string) {
	return s,y
}
//函数可以返回任意数量的返回值
func main() {
	a,b :=swap("hello","golang")
	fmt.Println(a,b)
}
```
- 命名返回值
``` 
package main

import "fmt"

func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum -x
	return 
}
//命名返回值  golang返回值可以命名 没有参数的return 返回结果当前值`


func main() {
	fmt.Println(split(17))
}
```
- 变量var  bool默认值 false 0
``` 
package main

import "fmt"

var cpp, python,php,java bool
//定义在func外面为全局变量
func main() {
	var i int
	//定义在func内为局部变量 
	fmt.Println(i,cpp,python,php,java)
}

```
- `:=` 简洁赋值语句
``` 
package main

import "fmt"

func main() {
	var i,j int = 1, 2
	c,python,java,php := true,true,false,"ok!"
	//在函数中可以使用 `:=` 简洁赋值语句在明确类型的地方可以用于替代var定义  BUT在函数体外必须以关键字开始 `:=`不能用于函数体外
	fmt.Println(i,j,c,python,java,php)
}

```
- 基础类型
> golang占位符 https://studygolang.com/articles/2644
``` 
package main
  
import (
        "fmt"
        "math/cmplx"
)

var (
        ToBe bool = false
        MaxInt uint64 = 1<<64-1
        z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
        const f = "%T(%v)\n"
        fmt.Printf(f,ToBe,ToBe)
        fmt.Printf(f,MaxInt,MaxInt)
        fmt.Printf(f,z,z)
}
// 别名 byte//uint8  rune //int32 代表一个Unicode码

//变量的定义可以打包在一个语法块中:

```


