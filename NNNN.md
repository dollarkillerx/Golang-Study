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

#### 比较
- 关于数组的比较,长度不同的数组不能比较 
  ```
  func TestCompareArray(t *testing.T) {
	a:=[...]int{1,2,3,4}
	//b:=[...]int{1,2,3,4,5}
	//c:=[...]int{1,2,3,4,5,6}
	d:=[...]int{1,2,3,4}
	//t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
	}
  ```
- 切面
  - 切片是一个结构体 包含ptr指针len总长度(可访问的总长度)cap容器长度
  - 当cap超出,会在内存new一个现在cap*2的长度的数组,然后吧现在的copy过去
  - 共享存储
	```
	year := []string{"Jan","Feb","Mar","Apr","May","Oct"}
	q2 := year[3:6]
	此为连续存储
	```
  - 声明
	```
	year := []string{}
	year = make([]int,2,6)
	```
### Map
- 声明
	```
	m := male[key]value{key:value}
	m2 := make(map[string]int,10/*Initial Capacity 内存空间 避免在此分配copy消耗更多时间*/)
	```
#### MAP与工厂模式
```
func TestMapWithFunValue(t *testing.T) {
	m:=map[int]func(op int)int{}
	m[0] = func(op int)int{return op}
	m[1] = func(op int)int {return op*op}
	m[2] = func(op int) int {return op*op*op}

	t.Log(m[0](2))
	t.Log(m[1](2))
	t.Log(m[2](2))
}
```
#### MAP实现set
- 元素唯一性
- 操作 正删改查 判断是否存在

### 字符串
- string 是数据类型,不是引用或者指针类型
- string是只读的byte slice,len函数可以它所包含的byte数
- string的byte数组可以存放任何数据
- 字符串转slice
	```
	s:="A,b,c"
	paras := strings.Split(s,",")
	t.Log(paras)
	```
- slice转str
	```
	t.Log(strings.Join(paras,"-"))
	```
- int转str	`s:=strconv.Itoa(10)`
- str转int `i:=strconv.Atoi("10")`
	```
	func TestConv(t *testing.T) {
	s:=strconv.Itoa(10)
	t.Logf("%T",s)
	if i,err :=strconv.Atoi("10");err==nil{
		t.Logf("%T",i)
	}
	}
	```
### 函数一等公民
- Golang都是值传递
  - 有些会觉得slice不是引用传递吗?哈哈你你们的slice的学习没有学好吧!slice是结构体pre是指针地址,你传了一个slice指针指向的地址还是同一个地址,就会有一种传递应用的错觉
  - 所有参数都是值传递,slice,map,channel会有传引用的错觉
  - 函数可以作为变量的值
  - 函数可以作为参数和返回值

### 特色的golang OOD
- Duck typing
```
type Employee struct {
	Id string
	Name string
	Age int
}

// 第一种定义方式在实现对应方法被调用时,实例的成员会进行值复制
func (e Employee) String() string {
	return fmt.Sprintf("ID:%s-Age:%d",e.Id,e.Name,e.Age)
}
// 通常情况下为避免内存拷贝,我们使用第二种定义方式
func (e *Employee) String() string {
	return fmt.Sprintf("ID:%s-Age:%d",e.Id,e.Name,e.Age)
}
```
#### 接口 duck typing
  ```
  type Programmer interface {
	  WriteHelloWOrld() Code
  }

  type GoProgrammer struct {

  }

  func (p *GoProgrammer) WriteHelloWorld() Code {
	  return "fmt.Println(\"Hello World\")"
  }
  ```
  - 接口为非入侵式,实现不依赖与接口定义
  - 所以接口的定义可以包含在接口使用者内
#### 扩展与复用
```
type Pet struct {

}

func (p *Pet) Speak()  {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ",host)
}

type Dog struct {
	p *Pet
}


Dog要去使用的话就需要d.p.Speak()
go提供了匿名嵌套类型
type Dog struct {
	Pet
}
d.Speak() 就行了
but 不支持访问儿子的数据 也不支持SLP
```
#### 空接口与断言
- 空接口可以代表任意类型
- 通过断言将空接口转换为指定类型
  - v,ok:=p.(int) //ok=true时为转换成功
  ```
	func DoSomething(p interface{}){
		if i,ok:=p.(int);ok{
			fmt.Println("int",i)
		}
		if i,ok:=p.(string);ok{
			fmt.Println("string",i)
		}
	}

	func TestEmptyInterfaceAssertion(t *testing.T)  {
		DoSomething("1616")
	}

	OR
	switch v:=p.(type) {
		case int:
			fmt.Println("this is int",v)
		case string:
			fmt.Println("this is string",v)
	}
  ```
- GO接口最佳实践
  - 倾向于使用小的接口定义,更多接口只包含以一个方法
  - 较大的接口定义,可以由多个小结构定义组合而成
  - 只以来必要功能的最小接口

### 错误机制
- 没有异常机制
- error类型实现类error接口
	```
	type error interface {
		Error() string
	}
	```
- 可以通过error.new来快速创建错误实例
	```
	errors.New("n must be in the range")
	```
- golang失败要求快速失败  (尽早失败,避免嵌套)
	```
	func GetFibonacci(n int) ([]int,error) {
		if n <0 || n>100 {
			return nil,errors.New("n should be in [0,100]")
		}
		fibList := []int{1,1}
		for i:=2;i<n;i++{
			fibList = append(fibList,fibList[i-2]+fibList[i-1])
		}
		return fibList,nil
	}

	func TestGetFibonacci(t *testing.T) {
		if v,err:=GetFibonacci(-10);err!=nil{
			t.Log("错误了")
		}else{
			t.Log(v)
		}
	}
	```
### Panic与Reciver
- panic用于不会恢复的错误
- panic退出前会执行defer指定的内容
- os.Exit 退出时不会调用defer指定的函数
- os.Exit 退出时不会输出当前调用栈信息
```
func TestPanVxExit(t *testing.T){
	defer func() {
		if err := recover();err!=nil{
			fmt.Println("恢复错误",err)
		}
	}()
	fmt.Println("Start")
	//os.Exit(15)
	panic(errors.New("Something wrong!"))
}
```
- 当心recover成为恶魔
	- 形成僵尸服务进程,导致health check失效
	- "Let it Crash!" 往往是我们恢复不确定性错误的最好方法
  
### 可复用的包
- 基本复用模块单元
  - 以首字母大写来表明可被包外代码访问
- 代码的package可以和所在目录不一致
- 同一个目录里的Go代码的package要保持一致
#### init方法
- 在main被执行前,所有依赖的package的init方法都会被执行
- 不同包的init函数按照包导入的依赖关系决定执行顺序
- 每个包可以有多个init函数
- 包的每个源文件也可以有多个init函数,这点比较特殊
  ```
  func init() {
		fmt.Println("this is init")
	}

	func init() {
		fmt.Println("this is init2")
	}
  ```	
- 获取远程库
  - go get -u ....
  - gopm 工具加速下载

### 依赖管理
- 随着Go1.5release版本发布,vendir目录被添加到除了GOPATH和GOROOT之外的依赖目录查找的解决方案
- 查询以来解决如下
  - 当前包下的vendor目录
  - 向上级目录查找,直到找到src下的vendor目录
  - 在GOPATH下查找依赖包
  - 在GOROOT下
- 推荐VGO 官方推荐 应该会在2.0或者g1.13后加入

### 协程
- 系统线程挂载P(go协程控制器)挂载协程队列
- 当当前协程队列某一个协程耗时非常长如何解决
  - p会每隔一段时间监控协程的数量,发现没有怎么变化会被当前协程终止放到队尾(状态会被保存)
  - 当某一个协程被系统中断了,需要等待了 未来提高并发p会移动到空闲的系统线程当中,继续执行队列
```
func TestGroutine(t *testing.T) {
	for i:=0;i<10;i++{
		go func (i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Microsecond * 50)
}
go 启动携程 (共享内存)
```
#### CSP 并发机制
csp vs actor
- 和Actor的直接通讯不同,CSP模式则是通过Channel进行通讯的,更松耦合一些
- Go中的channel是有容量现在并且独立于Grountine,而Erlang,Actor模式中的mailbox容量是无限的,接收进程也总是被动处理消息
- 