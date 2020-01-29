package main

import "fmt"
import "strings"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	sum := 0
	for _, name := range users{
		if strings.Contains(name, "e") || strings.Contains(name, "E"){
			distribution[name] = distribution[name] + 1
			sum ++
		}
		if strings.Contains(name, "i") || strings.Contains(name, "I"){
			distribution[name] = distribution[name] + 2
			sum = sum + 2
		}
		if strings.Contains(name, "o") || strings.Contains(name, "O"){
			distribution[name] = distribution[name] + 3
			sum = sum + 3
		}
		if strings.Contains(name, "u") || strings.Contains(name, "U"){
			distribution[name] = distribution[name] + 4
			sum = sum + 4
		}
	}
	return coins - sum
}

func main() {

	left := dispatchCoin()
	fmt.Println("剩下：", left)
	for k, v := range distribution {
		fmt.Println(k,v)
	}
}
