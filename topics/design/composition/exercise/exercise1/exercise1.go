package main

import "fmt"

// 示例为模拟一个公司管理工作流程.

type administrator interface {
	administrate(system string)
}

type developer interface {
	develop(system string)
}

type adminlist struct {
	list []administrator
}

func (l *adminlist) Enqueue(a administrator) {
	l.list = append(l.list, a)
}

func (l *adminlist) Dequeue() administrator {
	a := l.list[0]
	l.list = l.list[1:]

	return a
}

type devlist struct {
	list []developer
}

func (l *devlist) Enqueue(d developer) {
	l.list = append(l.list, d)
}

func (l *devlist) Dequeue() developer {
	d := l.list[0]
	l.list = l.list[1:]

	return d
}

// 定义具体实例
type sysadmin struct {
	name string
}

func (s *sysadmin) administrate(system string) {
	fmt.Printf("Name: %s, is managing %s \n", s.name, system)
}

type programmer struct {
	name string
}

func (p *programmer) develop(system string) {
	fmt.Printf("Name: %s, is coding %s \n", p.name, system)
}

type Company struct {
	administrator
	developer
}

func main() {

	// Create a variable named admins of type adminlist.
	var admins adminlist

	// Create a variable named devs of type devlist.
	var devs devlist

	// Enqueue a new sysadmin onto admins.
	admins.Enqueue(&sysadmin{name: "Li"})

	// Enqueue two new programmers onto devs.
	devs.Enqueue(&programmer{name: "Ye"})
	devs.Enqueue(&programmer{name: "Ming"})

	// Create a variable named cmp of type company, and initialize it by
	// hiring (dequeuing) an administrator from admins and a developer from devs.
	com := Company{
		administrator: admins.Dequeue(),
		developer:     devs.Dequeue(),
	}

	// Enqueue the company value on both lists since the company implements
	// each interface.
	admins.Enqueue(&com)
	devs.Enqueue(&com)

	// A set of tasks for administrators and developers to perform.
	tasks := []struct {
		needsAdmin bool
		system     string
	}{
		{needsAdmin: false, system: "xenia"},
		{needsAdmin: true, system: "pillar"},
		{needsAdmin: false, system: "omega"},
	}

	// Iterate over tasks.
	for _, task := range tasks {

		// Check if the task needs an administrator else use a developer.
		if task.needsAdmin {

			// Dequeue an administrator value from the admins list and
			// call the administrate method.
			ad := admins.Dequeue()
			ad.administrate(task.system)

			continue
		}

		// Dequeue a developer value from the devs list and
		// call the develop method.
		dr := devs.Dequeue()
		dr.develop(task.system)
	}
}
