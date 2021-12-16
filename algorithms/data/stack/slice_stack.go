package stack

// SliceStack 利用切片实现栈的数据结构。
// 栈的核心为后进先出（LIFO: Last in, First Out）的方式存储和删除元素
type SliceStack struct {
	stack []interface{}
}

// NewSliceStack 初始化新栈
func NewSliceStack() *SliceStack {
	return &SliceStack{
		stack: []interface{}{},
	}
}

func (s *SliceStack) Push(v interface{}) {
	s.stack = append(s.stack, v)
}

func (s *SliceStack) Pop() interface{} {
	if len(s.stack) == 0 {
		return nil
	}
	p := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return p
}

func (s *SliceStack) Peek() interface{} {
	if len(s.stack) == 0 {
		return nil
	}
	return s.stack[len(s.stack)-1]
}

func (s *SliceStack) Len() int {
	return len(s.stack)
}
