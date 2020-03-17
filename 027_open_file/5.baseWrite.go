package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./tmp.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file fiailed, err:", err)
		return
	}
	defer file.Close()
	str := "hello 世界\n"
	file.Write([]byte(str))
	file.WriteString("hello, tom")
}
