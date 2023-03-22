package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var Today = time.Now().Format("2006-01-02")

func Init(base_dir string) (err error) {
	fp := fmt.Sprintf("%v/logs/%v.log", base_dir, Today)
	logFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[eeo] ")
	return
}
