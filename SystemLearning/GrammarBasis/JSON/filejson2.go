package main

import (
	"encoding/json"
	"os"
)

type persio struct {
	Name string
	Age int
	Rmb float64
	Sex bool
	Hobby []string
}

func main() {
	filePath := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/JSON/config1.json"

	qianGe1 := persio{"于谦1", 50, 123.45, true, []string{"抽1烟", "喝酒", "烫头"}}
	qianGe2 := persio{"于谦2", 50, 123.45, true, []string{"抽2烟", "喝酒", "烫头"}}
	qianGe3 := persio{"于谦3", 50, 123.45, true, []string{"抽3烟", "喝酒", "烫头"}}
	p := make([]persio, 0)
	p = append(p,qianGe1)
	p = append(p,qianGe2)
	p = append(p,qianGe3)

	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	encoder := json.NewEncoder(file)
	err := encoder.Encode(&p)
	if err!=nil {
		panic(err.Error())
	}
}