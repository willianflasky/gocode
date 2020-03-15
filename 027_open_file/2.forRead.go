package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}

	defer file.Close()

	// 循环读取

	var content []byte
	var tmp = make([]byte, 128)

	for {
		n, err := file.Read(tmp)
		// 判断读完
		if err == io.EOF {
			fmt.Println("读取完了")
			break
		}
		// 判断读取错误
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}
