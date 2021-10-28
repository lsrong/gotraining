package effective_go_functions

import (
	"io"
	"os"
)

// Contents 读取文件里面的所有内容.
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	var result []byte
	buf := make([]byte, 100)
	for {
		// 分段读取内容
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(result), nil
}
