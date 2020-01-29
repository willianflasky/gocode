package main

import "fmt"

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)

	// 循环channel方法一
	for ret := range ch2 {
		fmt.Println(ret)
	}

}

func f1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
func f2(ch1 <-chan int, ch2 chan<- int) {
	// 循环channel方法二
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

