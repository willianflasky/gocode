package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {
	job := make(chan int64, 100)
	result := make(chan int64, 100)
	wg.Add(1)
	go createRand(job)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go readRand(job, result)
	}

	wg.Wait()
	close(result)
	for i := range result {
		fmt.Println(i)
	}
}

func createRand(job chan<- int64) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		job <- rand.Int63()
	}
	close(job)
}

func readRand(job <-chan int64, result chan<- int64) {
	defer wg.Done()
	for {
		value, ok := <-job
		if !ok {
			fmt.Println("没有取到值:", value, ok)
			break
		} else {
			mysum := randSum(value)
			result <- mysum
			fmt.Printf("随机值: %v 和: %v\n", value, mysum)
		}
	}
	fmt.Println("线程结束")
}

func randSum(x int64) int64 {
	var sum int64
	sum = 0
	for x > 0 {
		a := x % 10
		x = x / 10
		sum += a
	}
	return sum
}
