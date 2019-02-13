# Golang-Study
Golang Study
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