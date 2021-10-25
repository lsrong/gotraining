package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	filepath := "./reader.go.gz"
	f, e := os.Open(filepath)
	if e != nil {
		fmt.Printf("Failed to open file[%s]\n, error[%s]\n", filepath, e)
	}
	defer f.Close()
	reader, e := gzip.NewReader(f)
	if e != nil {
		fmt.Printf("Failed to create new  gzip reader [%s]\n", e)
	}
	var (
		content []byte
		buf     [128]byte
	)
	for {
		n, e := reader.Read(buf[:])
		if e == io.EOF {
			if n != 0 {
				content = append(content, buf[:n]...)
			}

			break
		}
		if e != nil {
			fmt.Printf("Failed to read file [%s]\n", e)
		}
		content = append(content, buf[n:]...)
	}

	fmt.Println(string(content))
}
