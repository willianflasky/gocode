package main

import (
	"fmt"
)

type Sayer interface {
	Say()
}

type Cat struct{ color string }
type Dog struct{ color string }
type Sheep struct{ color string }

func (this Cat) Say() {
	fmt.Println(this.color, "喵喵喵~")
}

func (this Dog) Say() {
	fmt.Println(this.color, "汪汪汪~")
}

func (this Sheep) Say() {
	fmt.Println(this.color, "咩咩咩~")
}

func main() {
	c := Cat{color: "black"}
	d := Dog{color: "yellow"}
	s := Sheep{color: "white"}

	general(c)
	general(d)
	general(s)
}

func general(s Sayer) {
	s.Say()
}
