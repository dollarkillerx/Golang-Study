package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonCotent := `{"age":15,"hobby":["抽烟","喝酒","烫头"],"name":"于谦","rmb":123.45,"sex":true}`

	bytes := []byte(jsonCotent) // 得到json字节
	dataMap := make(map[string]interface{})
	unmarshal := json.Unmarshal(bytes, &dataMap)
	if unmarshal != nil{
		panic(unmarshal.Error())
	}
	fmt.Println(dataMap)
}
