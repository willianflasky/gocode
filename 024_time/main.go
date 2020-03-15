package main

import (
	"fmt"
	"time"
)

func main() {
	//  1、时间类型
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	minute := now.Minute()
	second := now.Second()

	fmt.Println(now)
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)
	fmt.Println(minute)
	fmt.Println(second)

	// 2、时间戳
	ts1 := now.Unix()
	ts2 := now.UnixNano()
	fmt.Println(ts1)
	fmt.Println(ts2)

	// 3、timestamp->timeObj
	obj := time.Unix(ts1, 0)
	fmt.Println(obj)

	// 4、time.Duration 包定义的一个类型
	// time.Duration表示1纳秒，time.Second表示1秒
	n1 := time.Now()
	later := n1.Add(time.Hour)
	fmt.Println(later)

	sub := n1.Sub(later)
	eq := n1.Equal(later)
	before := n1.Before(now)
	after := n1.After(now)
	fmt.Println(sub)
	fmt.Println(eq)
	fmt.Println(before)
	fmt.Println(after)

	// 5、定时器
	// ticker := time.Tick(time.Second * 3)
	// for i := range ticker {
	// 	fmt.Println(i)
	// }

	// 6、格式化
	n3 := time.Now()
	fmt.Println(n3.Format("2006-01-02 15:04"))

	// 7、时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2020/03/14 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(n3))

}
