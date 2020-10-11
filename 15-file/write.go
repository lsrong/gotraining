package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteFile(filepath string, content string, flag int) (err error) {
	file, err := os.OpenFile(filepath, flag, 0755)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return
}

func IoutilWriteFile(filepath string, content string) error {
	err := ioutil.WriteFile(filepath, []byte(content), 0755)
	if err != nil {
		return err
	}
	return nil
}

func BufWriteFile(filepath string, content string, flag int) error {
	file, err := os.OpenFile(filepath, flag, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}
	err = writer.Flush()
	return err
}

func main() {
	filepath := "./test.dat"
	err := IoutilWriteFile(filepath, "ioutill!")
	if err != nil {
		fmt.Printf("Write file error[%s]", err)
	}

	err = WriteFile(filepath, "hello world!", os.O_WRONLY|os.O_CREATE|os.O_APPEND)
	if err != nil {
		fmt.Printf("Write file error[%s]", err)
	}

	err = BufWriteFile(filepath, "Bufio !", os.O_WRONLY|os.O_CREATE|os.O_APPEND)
	if err != nil {
		fmt.Print(err)
	}

}
