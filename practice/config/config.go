package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

/*
校验配置参数结构体
读取文件，逐行解析文件内容
分析结构体参数属性匹配配置项并赋值
*/
func UnMarshalFile(filepath string, config interface{}) error {
	// 校验配置参数
	typeInfo := reflect.TypeOf(config)
	if typeInfo.Kind() != reflect.Ptr {
		// 指针类型校验
		return errors.New("Please enter point args")
	}
	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		// 结构体校验
		return errors.New("Please enter struct args")
	}

	// 读取文件
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return errors.New("Failed to read ini file!")
	}
	//fmt.Println(string(content))
	lines := strings.Split(string(content), "\n")
	for index, line := range lines {
		// TODO 读取每一行内容，寻找配置节点 => 具体赋值
		fmt.Println(index)
		fmt.Println(line)
	}

	return nil
}
