package main

import (
	"fmt"
)

func main() {
	sm := make(map[string][]string, 10)
	// sm["中国"] = make([]string, 0, 2)
	key := "中国"

	value, ok := sm[key]
	if !ok {
		value = make([]string, 0, 2)
	} else {
		value = append(value, "广州")
	}
	value = append(value, "北京", "上海")
	sm[key] = value
	fmt.Println(sm)
}
