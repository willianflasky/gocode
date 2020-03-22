package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求：如果通知worker中的循环停下来;

var wg sync.WaitGroup

func worker() {
	defer wg.Done()

	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}

}
func main() {
	wg.Add(1)
	go worker()
	wg.Wait()
	fmt.Println("main done")
}
