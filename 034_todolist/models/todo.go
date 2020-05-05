package models

import (
	"todolist/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// CreateAToDo 创建todo
func CreateAToDo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

// GetAllToDo 获取ALL
func GetAllToDo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// GetAToDo  获取1个
func GetAToDo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateAToDo 更新
func UpdateAToDo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

// DeleteAToDo 删除
func DeleteAToDo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
