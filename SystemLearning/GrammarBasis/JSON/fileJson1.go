package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	dataMap := make(map[string]interface{})
	dataMap["name"] = "于谦"
	dataMap["age"] = 15
	dataMap["rmb"] = 123.45
	dataMap["sex"] = true
	dataMap["hobby"] = []string{"抽烟", "喝酒", "烫头"}

	filePath := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/JSON/config.json"
	dstFile,_ := os.OpenFile(filePath,os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0666)

	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(dataMap)
	if err != nil{
		panic(err.Error())
	}
	defer func() {
		fmt.Println("编码成功")
		dstFile.Close()
	}()
}