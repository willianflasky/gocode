package main

// go run main.go -name "tony" --age 28 --married=true -d=1h30m

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// os.Args 是一个[]string
	// if len(os.Args) > 0 {
	// 	for index, arg := range os.Args {
	// 		fmt.Printf("args[%v]%v\n", index, arg)
	// 	}
	// }

	// 脚本参数
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "tom", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	flag.Parse()
	fmt.Println(name, age, married, delay)
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())

}
