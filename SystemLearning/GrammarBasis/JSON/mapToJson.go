package main

import (
	"encoding/json"
	"fmt"
)

func mapToJson() string {
	dataMap := make(map[string]interface{})
	dataMap["name"] = "于谦"
	dataMap["age"] = 15
	dataMap["rmb"] = 123.45
	dataMap["sex"] = true
	dataMap["hobby"] = []string{"抽烟", "喝酒", "烫头"}

	bytes, e := json.Marshal(dataMap)
	if e != nil{
		panic(e.Error())
	}
	return string(bytes)
}

func main() {
	toJson := mapToJson()
	fmt.Println(toJson)
}