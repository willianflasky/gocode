package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("workder")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("default!")
		}
	}
	defer wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker 2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("default2")
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
	fmt.Println("main done")
}
