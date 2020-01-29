package main

import (
	"fmt"
	"os"
)

/*
	学员信息系统
		1、添加学生。
		2、编辑学生。
		3、展示所有学员信息。
*/

func main() {
	sm := newStudentMgr()
	for {
		// 需求分析
		//	1、打印系统菜单
		showMenu()
		//  2、等待用户选择执行的选项
		var input int
		fmt.Print("请输入序号>")
		fmt.Scanf("%d\n", &input)
		// fmt.Println("用户输入的是：", input)
		//  3、执行用户选择的动作

		switch input {
		case 1:
			// add student
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			// edit student
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			// show student
			sm.showStudent()
		case 4:
			os.Exit(0)
		}
	}

}

func showMenu() {
	fmt.Println("\033[31;1m欢迎来到学员信息管理系统\033[0m")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出系统")
}

func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入学员信息")
	fmt.Print("请输入学员ID：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员name：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员class：")
	fmt.Scanf("%s\n", &class)
	stu := newStudent(id, name, class)
	return stu
}
