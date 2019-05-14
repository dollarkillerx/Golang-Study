package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	filePath := "/home/dollarkiller/github/Golang-Study/SystemLearning/GrammarBasis/JSON/config.json"
	dataMap := make(map[string]interface{})

	file, _ := os.Open(filePath)

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&dataMap)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(dataMap)
	defer func() {
		file.Close()
	}()
}