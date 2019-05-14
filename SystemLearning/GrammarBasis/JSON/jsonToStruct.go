package main

import (
	"encoding/json"
	"fmt"
)

type persion struct {
	Name string
	Age int
	Rmb float64
	Sex bool
	Hobby []string
}

func main() {
	jsonData := `{"Name":"于谦","Age":50,"Rmb":123.45,"Sex":true,"Hobby":["抽烟","喝酒","烫头"]}`
	bytes := []byte(jsonData)
	var per persion
	err := json.Unmarshal(bytes,&per)
	if err!=nil {
		panic(err.Error())
	}
	fmt.Println(per)
}