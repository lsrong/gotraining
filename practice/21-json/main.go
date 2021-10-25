package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Result struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data []interface{} `json:"data"`
}

// json序列化并写入文件
func WriteJson(filepath string) error {
	var results []*Result
	for i := 0; i < 10; i++ {
		result := &Result{
			Code: rand.Int(),
			Msg:  fmt.Sprintf("Code:%d", rand.Int()),
		}
		results = append(results, result)
	}

	content, err := json.Marshal(results)
	if err != nil {
		return errors.New("Failed to marshal json ")
	}

	err = ioutil.WriteFile(filepath, content, 0755)
	if err != nil {
		return err
	}

	return nil
}

// 读取json
func ReadJson(filepath string) (*Result, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New("Failed to read json file ")
	}
	var result *Result
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, errors.New("Failed to unmarshal json ")
	}
	return result, nil
}

func main() {
	filepath := "/Users/lsrong/Work/Project/Go/src/github.com/LearningGolang/21-json/demo.json"
	result, err := ReadJson(filepath)
	if err != nil {
		fmt.Printf("Failed to unmarshal json %v \n", err)
		return
	}
	fmt.Println(result)

	writeFile := "/Users/lsrong/Work/Project/Go/src/github.com/LearningGolang/21-json/write.json"
	err = WriteJson(writeFile)
	if err != nil {
		fmt.Printf("Failed to write json %v \n", err)
		return
	}
}
