package main

import "fmt"

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
		for _, char := range name{
			switch char {
			case 'e', 'E':
				distribution[name] = distribution[name] + 1
				sum +=1	
			case 'i', 'I':
				distribution[name] = distribution[name] + 2
				sum +=2
			case 'o', 'O':
				distribution[name] = distribution[name] + 3
				sum +=3
			case 'u', 'U':
				distribution[name] = distribution[name] + 4
				sum +=4
			}
		}
	}
	return coins - sum
}

func main(){
	s := dispatchCoin()
	fmt.Println(s)
	for k, v := range distribution {
		fmt.Println(k,v)
	}
}
