package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}


// 倒三角
// func main() {
// 	for i := 1; i < 10; i++ {
// 		for j := i; j < 10; j++ {
// 			fmt.Printf("%d * %d = %d\t", j, i, i*j)
// 		}
// 		fmt.Println()
// 	}
// }
