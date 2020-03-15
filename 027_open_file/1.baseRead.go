package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1、基本读取文件
	file, err := os.Open("./go.mod")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()

	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}

	fmt.Printf("读取字节数据：%v\n", n)
	fmt.Println(string(tmp[:n]))
}
