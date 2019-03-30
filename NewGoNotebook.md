重拾GOlang
===

### 基础
- 变量定义
``` 
var a int
var s string
var a,b int = 3,4
a,b := 3,4
fmt.Printf("%d %q\n",a,s)
```
> 在函数外部定义不能省略 var
- 常量与枚举
``` 
const filename = 'hello.gio'

const (
    cpp = 0
    java = 1
    python = 3
    golang = 3
)
fmt.Print(cpp,java,python,golang)

const (
    cpp = iota
    _
    python
    golang
    php
)
iota 自增

const (
    b = 1 << (10* iota)
    kb
    mb
)
iota 种子
```
- 条件语句 if
``` 
    if contents, err := ioutil.ReadFile(filename); err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}
```
> if 条件可以赋值  变量作用域就在这个if语句里
- 条件语句 switch
``` 
func eval(a,b int,op string) int  {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsc:" + op)
	}
	return result
}

func main()  {
	fmt.Println(eval(15,75,"+"))
}
```
or
``` 
func grade(score int) string {
    switch {
    case score < 60:
        return "f"
    case score > 80:
        return 'a'
    default:
        return 'a'
        return 'a'
    }
}
```
> switch会自动break 除非使用fallthrough
- if 循环
``` 
sum := 0
for i :=1;i<=100;i++ {
    sum += i
}
```
- 函数
``` 
func apply(op func(int,int) int,a,b int) int {
	fmt.Printf("Calling function %s with args" + "(%d,%d)",opName,a,b)
	return op(a,b)
}

func sum(numbers ...int) int {
    s := 0
    for i := range numbers {
        s += numbers[i]
    }
}
```
- 指针  golang默认都是值传递
``` 
var a = 2
var pa *int = &a
*pa = 3
fmt.Println(a)
```

``` 
func huan(a,b *int)  {
	*a,*b = *b,*a
}

func main() {
	a,b := 5,6
	huan(&a,&b)
	fmt.Printf("a:%d,b:%d",a,b)
}
```
- 数组
``` 
func main()  {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	var grid [4][5]int
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(grid)

	arr3 :=[...]int{2,4,5,9,8}
	for i:=0;i<len(arr3);i++{
		fmt.Println(arr3[i])
	}

	for k,v := range arr3 {
		fmt.Println(k,v)
	}
}
``` 
- 切片 前闭后开
```  
func main()  {
	arr :=[...]int{0,1,2,3,4,5,6,7}
	s := arr[2:6]
	fmt.Println(s)
}
```
- defer栈
``` 
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```
> 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
- 结构体struct
``` 
type Me struct {
    X int
    Y int
}

func main() {
   v := Me{1,2}
   fmt.Println(v.X)
   v.X = 4
   fmt.Println(v.X)
}
```
- 结构体指针
``` 
type Me struct {
    X int
    Y int
}

func main() {
    v := Me{1,5}
    p :=&v
    p.X = 1e9
    fmt.Println(v)
}
```
- 结构体文法
``` 
var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)
```
- 数组
``` 
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```
- 切片
> 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。
``` 
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```
> 切片并不存储任何数据，它只是描述了底层数组中的一段。
更改切片的元素会修改其底层数组中对应的元素。
与它共享底层数组的切片都会观测到这些修改。
- len(s)切片的长度 和 cap(s)切片的容量
- make创建切片
``` 
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
``` 
- 向切片追加元素
``` 
func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```
> func append(s []T, vs ...T) []T
  append 的第一个参数 s 是一个元素类型为 T 的切片，其余类型为 T 的值将会追加到该切片的末尾。
  append 的结果是一个包含原切片所有元素加上新添加元素的切片。
  当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。
- Range
``` 
    for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	
	for i := range pow {
        pow[i] = 1 << uint(i) // == 2**i
    }
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
```
- 映射
>映射将键映射到值。
 映射的零值为 nil 。nil 映射既没有键，也不能添加键
 make 函数会返回给定类型的映射，并将其初始化备用。
``` 
 dict1 := make(map[string]int) //空映射，等同于dict1 := map[string]int{}  key[string] value [int]
 
 type Vertex struct {
 	Lat, Long float64
 }
 
 var m map[string]Vertex
 
 func main() {
 	m = make(map[string]Vertex)
 	m["Bell Labs"] = Vertex{
 		40.68433, -74.39967,
 	}
 	fmt.Println(m["Bell Labs"])
 }
 
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```
- 修改映射
``` 
在映射 m 中插入或修改元素：

m[key] = elem
获取元素：

elem = m[key]
删除元素：

delete(m, key)
通过双赋值检测某个键是否存在：

elem, ok = m[key]
若 key 在 m 中，ok 为 true ；否则，ok 为 false。

若 key 不在映射中，那么 elem 是该映射元素类型的零值。

同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。

注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：

elem, ok := m[key]
```
- 函数的闭包
``` 
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```
- 方法  (方法只是个带接收者参数的函数)
``` 
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

Go 没有类。不过你可以为结构体类型定义方法。

方法就是一类带特殊的 接收者 参数的函数。

方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
```
### Beego构建web
- 前面写的好累啊,直接在框架里面学习算了
#### 准备工作
- `go env` 查看当前go 环境
 - 主要关注
    - go root  go的安装目录
    - go path  go的包  工作目录
        - src 工程源代码
        - bin 编译生成可执行文件
        - pkg 包文件
- Beego 环境搭建
    - `go get -u github.com/astaxie/beego`
    - bee工具包
        - beego项目创建
        - 热编译
        - 开发测试
        - 部署
- 特点
    - 快速才发
    - MVC架构
    - 文档齐全,社区活跃
- 架构及其原理
    - cache
        - 文件 内存 me 推荐redis
    - config
        - 推荐json
    - context
        - request
        - response
    - httplibs
        - get post put delete head
        - 支持https
        - 支持超时设置
        - 文件上传
    - logs
        - 异步输出
        - 支持多种引擎
    - orm
    - session
    - toolbox
        - 监控
        - 定时任务
- bee工具应用
    - `bee new :新建项目结构`
    - `bee run :自动编译部署`
    - `bee generate :自动生成代码`
``` 
bee new dollarkiller
```
目录结构
``` 
├── conf 配置
│   └── app.conf
├── controllers 控制器
│   └── default.go  默认控制器
├── dollarkiller 二进制文件 bee run 执行后的结果
├── main.go 文件入库
├── models model层
├── routers 路由
│   └── router.go
├── static 静态文件
│   ├── css
│   ├── img
│   └── js
│       └── reload.min.js
├── tests 测试代码
│   └── default_test.go
└── views 视图
    └── index.tpl
```
- 获取数据 绑定参数 渲染模板
``` 
    name := c.GetString("name")
	c.Data["name"] = name
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
```
```
bee api [appname] [-tables=""] [-driver=mysql] [-conn="root:<password>@tcp(127.0.0.1:3306)/test"] 
bee generate scaffold user -fields="id:int64,name:string,gender:int,age:int" -driver=mysql -conn="beego:@tcp(192.168.31.89:3306)/beego";
```