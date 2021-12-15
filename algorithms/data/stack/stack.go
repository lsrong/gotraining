package stack

import (
	"errors"
	"github.com/learning_golang/algorithms/data/list"
)

// Stack 栈的数据结构特点： 按照后进先出（LIFO: Last in, First Out）的方式存储和删除元素。
// 用单向链表实现的栈数据结构
type Stack struct {
	list list.List
}

// NewStack 初始化新的栈实例
func NewStack() *Stack {
	var l list.List

	return &Stack{l}
}

// Push 将新数据放入栈中（压栈）
func (s *Stack) Push(data interface{}) int {
	s.list.InsertAfter(data, nil)

	return s.Len()
}

// Pop 从栈顶淡出一个结点（出栈）
func (s *Stack) Pop() (interface{}, error) {
	n, err := s.list.RemoveAfter(nil)
	if err != nil {
		return nil, errors.New("stack is empty")
	}

	return n.Data(), nil
}

// Peek 获取栈顶元素
func (s *Stack) Peek() interface{} {
	if head := s.list.Head(); head != nil {
		return head.Data()
	}
	return nil
}

func (s *Stack) Len() int {
	return s.list.Len()
}
