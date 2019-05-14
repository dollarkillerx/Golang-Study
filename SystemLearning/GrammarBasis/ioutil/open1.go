package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func simpleOpen()  {
	file,error := os.Open("/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/TestFile.test")
	if error == nil {
		fmt.Println("打开文件")
	}else{
		panic(error.Error())
	}
	defer func() {
		file.Close()
		fmt.Println("文件关闭")
	}()
	fmt.Println("拿到一个文件",file)

	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println()
			fmt.Println("文件读完了")
			break
		}else if err != nil {
			panic(err.Error())
		}
		fmt.Println(str)
	}

}

func intermediateOpen()  {
	//地址,方式,权限
	file,err := os.OpenFile("/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/TestFile.test",os.O_RDONLY,0666)
	if err!=nil {
		panic(err.Error())
	}
	defer func() {
		file.Close()
		fmt.Printf("文件close成功")
	}()

	// 创建读写缓冲区
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF { // err == io.EOF 就读完了
			fmt.Println("============================")
			fmt.Println("文件到末尾了")
			break
		}
		fmt.Printf(str)
	}
}


func ioutilRead() {
	// ioutil 自己就执行了打开和关闭
	byt,err :=ioutil.ReadFile("/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/TestFile.test")
	if err !=nil {
		panic(err.Error())
	}
	content := string(byt)
	fmt.Println(content)
}

func ioReadCreate() {
	fileDir := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/TestFile0.test"
	file,err := os.OpenFile(fileDir,os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0666)// os.O_TRUNC 覆盖
	if err!=nil {
		panic(err)
	}
	defer func() {
		file.Close()
	}()
	write := bufio.NewWriter(file)
	write.WriteString("鸣神の　少しとよみて　さし昙り　雨も降らんか　君を留めん\n")
	write.WriteString("鸣神の　少しとよみて　降らずとも　我は止まらん　妹し留めば\n")
	write.WriteString("隐约雷鸣，阴霾天空，但盼风雨来，能留君在此。\n")
	write.WriteString("隐约雷鸣，阴霾天空，即使天无雨，吾亦留在处。\n")
	write.Flush()// 清空缓冲区(立刻将缓冲区剩余的数据写入)
}

func simpleWrite() {
	readFile := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/TestFile0.test"
	read,err := ioutil.ReadFile(readFile)
	if err != nil {
		panic(err.Error())
	}
	writeFile := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/TestFile1.test"
	read =[]byte(read)
	err = ioutil.WriteFile(writeFile,read,0666)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("写出成功")
}

func fileExist(file string) bool {
	_,err := os.Stat(file)
	if err!=nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err.Error())
		}
	}
	return true
}

func simpleCopy() {
	filePath := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/timg.jpeg"
	bytes,err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}
	copyPath := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/timeCopy.jpeg"
	fmt.Printf("%v,%T",bytes,bytes)
	err = ioutil.WriteFile(copyPath,bytes,0666)
	if err!=nil {
		panic(err.Error())
	}
}

func ioCopy() {
	filePath := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/timg.jpeg"
	writePath := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/timgCc.jpeg"
	readFile,err := os.OpenFile(filePath,os.O_RDONLY,0666)
	writeFile,err := os.OpenFile(writePath,os.O_CREATE|os.O_WRONLY,0666)
	wriften,err:= io.Copy(writeFile,readFile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("文件拷贝成功",wriften)
}

func bufferCopy()  {
	// 使用1k缓冲区配合缓冲读写器进行拷贝

	filePath := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/timg.jpeg"
	readFile,err := os.OpenFile(filePath,os.O_RDONLY,0666)
	if err!=nil {
		panic(err.Error())
	}
	read := bufio.NewReader(readFile)

	OutPath := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/NNN.jpeg"
	writeFile,err := os.OpenFile(OutPath,os.O_CREATE|os.O_WRONLY,0666)
	write := bufio.NewWriter(writeFile)
	buffer := make([]byte,1024)
	for {
		// 创建缓冲区
		_,err := read.Read(buffer)
		if err == io.EOF {
			break
		}
		write.Write(buffer)
	}

	defer func() {
		write.Flush()
		readFile.Close()
		writeFile.Close()
	}()

}


func main() {
	//simpleOpen()
	//intermediateOpen()
	//ioutilRead()
	//ioReadCreate()
	//simpleWrite()
	//file := "/home/dollasrkiller/github/Golang-Study/SystemLearning/GrammarBasis/ioutil/TestFile0.test"
	//if fileExist(file) {
	//	fmt.Println("文件存在")
	//}else{
	//	fmt.Println("文件不存在")
	//}
	//simpleCopy()
	//ioCopy()
	bufferCopy()
}