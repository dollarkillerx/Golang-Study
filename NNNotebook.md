# Go 慢慢来

### slice
slice 可向后扩展 是数组的视图
len() slice总长度
cap() 不可超越数组长度
添加元素如果超越cap,系统会重新分配更大的底层数组
`s2 := make([]int,16,25)`
len=16,切片的长度
cap=25切片的容量
```
	s2 := make([]int,16,25)
	printSlice(s2)

	fmt.Println("Copying slice")
	copy(s2,s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3],s2[4:]...)

	printSlice(s2)

	fmt.Println("Popping from front")
	s2 = s2[1:]
	printSlice(s2)
	fmt.Println("Popping from back")
	s2 = s2[:(len(s2)-1)]
	printSlice(s2)
```

### Map
```
m := map[keyType]valueType {
    "key":"value",
    "key":"value",
    "key":"value",
    "key":"value",
}


m := map[string]string {
		"name":"dollarkiller",
		"course":"golang",
		"site":"imooc",
		"quality":"notbad",
	}
	fmt.Println(m["name"])
	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int // m2 = nil
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for i:=range m{
		fmt.Println(i," : ",m[i])
	}
	fmt.Println("Getting values")
	fmt.Println()
	if 	couresName,ok := m["cours e"];ok{
		fmt.Println(couresName)
	}else{
		fmt.Println("k is empty")
	}

	a,_ := returnTest(1,5)
	fmt.Println(a)
	fmt.Println("Delete values")
	delete(m,"name")
	fmt.Println(m)
```
- 使用range遍历key,或者遍历key,value
- 不保证遍历顺序,如需顺序,需手动对key排序
- 使用len获得元素个数
- 实例:寻找最长不含有重复字符的子串
```
package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i,ch:=range []byte(s){
		fmt.Println(ch)
		fmt.Println(lastOccurred)
		if lastI,ok := lastOccurred[ch];ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 >maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main()  {
	fmt.Println(lengthOfNonRepeatingSubStr("asdrgerg"))
}
```
#### rune
- 使用range遍历pos,rune对
- 使用utf8.RuneCountInString获得字符数量
- len获得字节的长度
- []byte获得字节
- []rune 获得rune 会重新开辟一个数组 放入rune类型
- 其他操作strings.func()
  - 返回在str出现的次数
  ```
    s = "12314564"
	fmt.Println(strings.Count(s,"1"))
  ```
  - Fields
#### 结构体和方法
```
type treeNode struct {
	value int
	left,right *treeNode
}

func main() {
	var root treeNode
	fmt.Println(root) //{0 <nil> <nil>}
	root = treeNode{3,&treeNode{},&treeNode{}}
	fmt.Println(root)//{3 0xc00000c060 0xc00000c080}
	root = treeNode{left:&treeNode{}}
	fmt.Println(root)//{0 0xc00000c060 <nil>}
	root.right = &treeNode{value:3}
	fmt.Println(root)//{0 0xc00000c0e0 0xc00000c120}
	root.left.right = new(treeNode) //new 返回地址
	fmt.Println(root.left.right)//&{0 <nil> <nil>}

	nodes := []treeNode{
		{value:3},
		{},
		{6,nil,&root},
	}
	fmt.Println(nodes)//[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc000090020}]

	fmt.Println()
	root.print(5)
}


func createNode(value int) *treeNode {
	// 工厂函数控制struct的创建
	return &treeNode{value:value}
	// 重点: 局部遍历也可以返回给外面用
		// 创建在堆栈不需要知道
		// 如果局部变量取地址返回使用就会分配在堆上  参与垃圾回收
		// 反之没有使用就会分配到栈上
}

// (node treeNode)为接受者 为结构定义方法
func (node treeNode) print(a int) {
	fmt.Println(node," ",a)
}

// (node treeNode)为接受者 为结构定义方法
func (node *treeNode) print(a int) {
	node.value = 5
	fmt.Println(node," ",a)
}

// 注意了:golang默认说有都是值传递
```
- 值接受者VS指针接受者
  - 要改变内容使用指针接受者
  - 结构过大也要考虑指针接受者
  - 一致性:如果有指针接受者,最好指针接受者
#### 封装
- CanmelCase
- 首字母大写public 小写private
- 针对package
- main包包含一个main函数
- 结构定义方法必须放在同一个包内  可以是不同的文件
#### Package 如果扩充系统类型或者别人的类型
- 定义别名
- 使用组合


### 接口 duck typing
```
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}
```
接口是隐式的,实现者只要完成接口的方法就实现了接口 dock typing
- 接口变量自带指针
- 接口变量通用采用值传递,几乎不需要使用接口的指针
- 指针接受者实现只能以指针方式使用;值接受者两者都可以
#### 查看接口遍历
- 标示任何类型: interface{}
- Type Assertion
- Type Switch
```
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q,v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool  {
	return len(*q) == 0
}

	q := queue.Queue{1}
	q.Push(2)
	q.Push("你好")
	q.Push(false)
	fmt.Println(q)
```
#### 函数与闭包
- 函数是一等工贸:参数,遍历,返回值都可以是函数
- 高阶函数
- 函数->闭包
- 正统的函数式编程
  - 不可变性:不能有状态,只有常量和函数
  - 函数只能有一个参数
```
func adder() func(int) int {
	sum := 0 // 自由变量  会被保存下来
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	a := adder()
	for i:=0;i<10;i++  {
		fmt.Println(a(i))
	}
}
```

#### 资源管理与出错处理
- defer调用  确保调用在函数结束时发生
```
func writeFile(filename string) {
	if file,err := os.Create(filename);err!=nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i:=0;i<20;i++ {
		fmt.Fprintln(writer,f())
	}
}
```
- 参数咋idefer语句时计算
- defer列表为栈