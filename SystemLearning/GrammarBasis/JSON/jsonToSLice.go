package main

import (
	"encoding/json"
	"fmt"
)

var jsonData =  `[{"age":15,"hobby":["抽烟","喝酒","烫头"],"name":"于谦","rmb":123.45,"sex":true},{"age":15,"hobby":["抽烟","喝酒","烫头"],"name":"于谦","rmb":123.45,"sex":true},{"age":15,"hobby":["抽烟","喝酒","烫头"],"name":"于谦","rmb":123.45,"sex":true},{"age":15,"hobby":["抽烟","喝酒","烫头"],"name":"于谦","rmb":123.45,"sex":true}]`

func main() {
	dataSlice := make([]map[string]interface{},0)
	bytes := []byte(jsonData)
	unmarshal := json.Unmarshal(bytes, &dataSlice)
	if unmarshal != nil{
		panic(unmarshal.Error())
	}
	fmt.Println(dataSlice)
}