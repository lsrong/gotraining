package main

import (
	"fmt"
)

type Manager struct {
	students []*Student
}

// 查询
func (m *Manager) ShowStudent() {
	if len(m.students) == 0 {
		fmt.Println("暂无学生")
		return
	}
	for i, v := range m.students {
		fmt.Printf("学生【%d】:%v \n", i, v)
	}
}

// 添加
func (m *Manager) AddStudent(stu *Student) {
	for i, v := range m.students {
		if v.Name == stu.Name {
			m.students[i] = stu
			return
		}
	}
	m.students = append(m.students, stu)
	fmt.Println("添加成功！")
}

// 修改
func (m *Manager) EditStudent(stu *Student) {
	for i, v := range m.students {
		if v.Name == stu.Name {
			m.students[i] = stu
			fmt.Println("修改成功！")
			return
		}
	}
	fmt.Println("学生没有找到！")
}

// 删除
func (m *Manager) deleteStudent(name string) {
	for i, v := range m.students {
		if v.Name == name {
			m.students = append(m.students[:i], m.students[i+1:]...)
			fmt.Println("删除成功")
			return
		}
	}
	fmt.Println("学生没有找到！")
}
