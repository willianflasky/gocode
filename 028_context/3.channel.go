package main


import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(exitChannel chan int){
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select{
		case <- exitChannel:
			break LOOP
		default:
			fmt.Println("default here!")
		}
	}
	wg.Done()
}

func main(){
	var exitChannel = make(chan int, 10)
	wg.Add(1)
	go worker(exitChannel)
	exitChannel <- 1
	time.Sleep(time.Second*3)
	close(exitChannel)
	wg.Wait()
	fmt.Println("main done")
}