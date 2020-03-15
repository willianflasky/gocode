package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	logFile, err := os.OpenFile("./main.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file field: ", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[小聪明]")
}

func main() {
	log.Println("这是一条很普通的日志")
	log.Fatalln("我是一个fatal的小日志")
	log.Panicln("我是一个panic")

	// 方法二：
	// logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	// logger.Println("这是自定义的logger记录的日志。")

	// Go内置的log库功能有限，例如无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等。
}
