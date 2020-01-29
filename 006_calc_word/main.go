package main

import (
	"fmt"
	"strings"
)

//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
func main() {
	str := "how do you do think you about"
	strSlice := strings.Split(str, " ")
	fmt.Println(strSlice)

	countMap := make(map[string]int, 10)
	for _, key := range strSlice {
		_, isReal := countMap[key]
		if !isReal {
			countMap[key] = 1
		} else {
			countMap[key] += 1
		}
	}
	fmt.Println(countMap)
}
