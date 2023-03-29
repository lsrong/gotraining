package student

import "fmt"

type Manager struct {
	students []*Student
}

// Show 展示学员信息
func (m *Manager) Show() {
	if len(m.students) == 0 {
		fmt.Println("还未录入学生")
		return
	}

	for i, v := range m.students {
		fmt.Printf("学生【%d】:%v \n", i, v)
	}
}

// Edit 编辑学员
func (m *Manager) Edit(std *Student) {
	for i, v := range m.students {
		if v.Name == std.Name {
			m.students[i] = std
			return
		}
	}
	fmt.Printf("学员【%v】没有找到！\n", std)
}

// Add 添加学员
func (m *Manager) Add(std *Student) {
	// 对比是否学生是否存在，存在则替换
	for i, v := range m.students {
		if v.Name == std.Name {
			m.students[i] = std
			return
		}
	}

	m.students = append(m.students, std)
}

// Delete 删除学员
func (m *Manager) Delete(name string) {
	for i, v := range m.students {
		if v.Name == name {
			m.students = append(m.students[:i], m.students[i+1:]...)
			fmt.Println("删除成功！")
			return
		}
	}

	fmt.Printf("学生【%s】没有找到", name)
}
