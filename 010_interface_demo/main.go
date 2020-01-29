package main

import "fmt"

// People interface
type People interface {
	Speak(string) string
}

// Student struct
type Student struct{}

// Speak method
func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	//1. 当Speak参数要指针时，这里必须给Student指针。 2. 当Speak参数要Student类型时，这里传Student指针和类型都可以。
	var peo People = Student{}
	think := "sb"
	fmt.Println(peo.Speak(think))
}

