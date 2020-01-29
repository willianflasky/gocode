package main

import "fmt"

func main() {
	for i := 200; i < 1000; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf("%d是质数\n", i)
		}
	}
}

