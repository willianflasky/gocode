package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMgr struct {
	allStudent []*student
}

func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudent: make([]*student, 0, 100),
	}
}

func (s *studentMgr) addStudent(newStu *student) {
	s.allStudent = append(s.allStudent, newStu)
}

func (s *studentMgr) modifyStudent(newStu *student) {
	for k, v := range s.allStudent {
		if newStu.id == v.id {
			s.allStudent[k] = newStu
			return
		}
	}
	fmt.Printf("没有找到这个学生:%d\n", newStu.id)

}

func (s *studentMgr) showStudent() {
	for _, v := range s.allStudent {
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", v.id, v.name, v.class)
	}
}
