package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit bool

func worker() {
	defer wg.Done()
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit{
			break
		}
	}
}
func main(){
	wg.Add(1)
	go worker()
	exit = true
	wg.Wait()
	fmt.Println("main done")
}