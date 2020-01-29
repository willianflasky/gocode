package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = [...]int{3, 7, 8, 9, 1}
	// 升序
	sort.Ints(a[:])
	fmt.Println(a)
	// 降序
	b := a[:]
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	fmt.Println(b)
}
