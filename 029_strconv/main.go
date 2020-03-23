package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1.字符串->数字
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("cat't convert to int")
	} else {
		fmt.Printf("type: %T value: %#v \n", i1, i1)
	}
	// 2.数字->字符串
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type: %T value: %#v\n", s2, s2)

	// Parse系列函数
	// 3.转布尔：1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	fmt.Println(strconv.ParseBool("F"))

	// 4. 字符串转浮点
	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("%T %v\n", f, err)

	// 5. 返回字符串类型是int64/uint64
	i, err1 := strconv.ParseInt("-2", 10, 64)
	u, err2 := strconv.ParseUint("2", 10, 64)
	fmt.Printf("%T, %v\n", i, err1)
	fmt.Printf("%T, %v\n", u, err2)

	// Format系列函数
	// float/int/unit转成字符串
	f1 := strconv.FormatBool(true)
	f2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	f3 := strconv.FormatInt(-2, 16)
	f4 := strconv.FormatUint(2, 16)
	fmt.Printf("%T %#v\n", f1, f1)
	fmt.Printf("%T %#v\n", f2, f2)
	fmt.Printf("%T %#v\n", f3, f3)
	fmt.Printf("%T %#v\n", f4, f4)

	f5 := strconv.CanBackquote("xx\n")
	fmt.Println(f5)

	f6 := strconv.IsPrint('中')
	fmt.Println(f6)

}
