package main

import (
	"fmt"
	"reflect"
)

func main() {
	stu1 := student{
		Name:  "张三",
		Score: 18,
	}
	printMethod(stu1)
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(v.FieldByName("Name"))
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", v.Method(i).Type())
		var args = []reflect.Value{v.FieldByName("Name"), v.FieldByName("Score")}
		v.Method(i).Call(args)
	}
}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s student) Study(name string, score int) string {
	msg := "好好学习天天向上"
	fmt.Println(msg, name, score)
	return msg
}

func (s student) Sleep(name string, score int) string {
	msg := "好好睡觉"
	fmt.Println(msg, name, score)
	return msg
}

