package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	for id := 1; id <= 3; id++ {
		go worker(id, jobs, result)
	}
	// 给5个任务
	for job := 1; job <= 5; job++ {
		jobs <- job
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-result
	}
}

func worker(id int, jobs chan int, result chan int) {
	for job := range jobs {
		fmt.Printf("工作者: %d start 任务号: %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("工作者: %d done 任务号: %d\n", id, job)
		result <- job * 2
	}
}

