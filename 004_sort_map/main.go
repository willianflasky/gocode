package main


import "fmt"
import "sort"

func main(){
	// 定义map
	mymap := make(map[string]int, 8)
	mymap["d"] = 1
	mymap["a"] = 2
	mymap["b"] = 3
	mymap["c"] = 4

	// 定义切片
	var keys = make([]string, 0, 100)
	for key := range(mymap) {
		keys = append(keys, key)
	}
	// 排序前
	fmt.Println(keys)
	sort.Strings(keys)
	// 排序后
	fmt.Println(keys)

}
