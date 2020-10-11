package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 直接读取文件内容
func ReadFile(filepath string) (content string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}

	defer file.Close()
	var (
		n   []byte
		buf [128]byte
	)
	for {
		line, err := file.Read(buf[:])
		// 是否终止
		if err == io.EOF {
			break
		}

		// 错误
		if err != nil {
			return "", err
		}

		// 追加
		n = append(n, buf[:line]...)
	}
	content = string(n)

	return
}

// bufio读取文件
func BufReadFile(filepath string) (content string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		content = fmt.Sprintf("%s%s", content, line)
	}
	return
}

// ioutil 读取整个文件
func IoutilReadFile(filepath string) (content string, err error) {
	n, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	content = string(n)

	return
}

func main() {
	filepath := "./file.json"

	// file.Read
	fmt.Println("ReadFile")
	content, err := ReadFile(filepath)
	if err != nil {
		fmt.Printf("Failed to read file[%s], error[%s]", filepath, err)
	}
	fmt.Println(content)

	// bufio.NewReader(file)  reader.ReadString('\n')
	fmt.Println("BufReadFile")
	bufContent, err := BufReadFile(filepath)
	if err != nil {
		fmt.Printf("Failed to read file[%s], error[%s]", filepath, err)
	}
	fmt.Println(bufContent)

	// ioutil
	fmt.Println("IoutilReadFile")
	ioutilContent, err := IoutilReadFile(filepath)
	if err != nil {
		fmt.Printf("Failed to read file[%s], error[%s]", filepath, err)
	}
	fmt.Println(ioutilContent)
}
