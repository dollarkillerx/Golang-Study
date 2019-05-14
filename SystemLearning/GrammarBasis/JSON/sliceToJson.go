package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	dataSlice := make([]map[string]interface{},0)

	dataMap1 := make(map[string]interface{})
	dataMap1["name"] = "于谦"
	dataMap1["age"] = 15
	dataMap1["rmb"] = 123.45
	dataMap1["sex"] = true
	dataMap1["hobby"] = []string{"抽烟", "喝酒", "烫头"}

	dataMap2 := make(map[string]interface{})
	dataMap2["name"] = "于谦"
	dataMap2["age"] = 15
	dataMap2["rmb"] = 123.45
	dataMap2["sex"] = true
	dataMap2["hobby"] = []string{"抽烟", "喝酒", "烫头"}

	dataMap3 := make(map[string]interface{})
	dataMap3["name"] = "于谦"
	dataMap3["age"] = 15
	dataMap3["rmb"] = 123.45
	dataMap3["sex"] = true
	dataMap3["hobby"] = []string{"抽烟", "喝酒", "烫头"}

	dataMap4 := make(map[string]interface{})
	dataMap4["name"] = "于谦"
	dataMap4["age"] = 15
	dataMap4["rmb"] = 123.45
	dataMap4["sex"] = true
	dataMap4["hobby"] = []string{"抽烟", "喝酒", "烫头"}

	dataSlice = append(dataSlice,dataMap1)
	dataSlice = append(dataSlice,dataMap2)
	dataSlice = append(dataSlice,dataMap3)
	dataSlice = append(dataSlice,dataMap4)

	bytes, e := json.Marshal(dataSlice)
	if e!=nil {
		panic(e.Error())
	}
	fmt.Println(string(bytes))
}
