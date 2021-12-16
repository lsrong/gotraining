package stack

// 将持续更新力扣关于栈的练习...

// IsValidSymbol 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 示例：
// 1. s = "()"， true ;
// 2. s = "()[]{}"， true ;
// 3. s = "(]"， false ;
// 4. s = "([)]"， false ;
// 5. s = "{[]}"， true ;
func IsValidSymbol(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := NewStack()
	symbols := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	for _, b := range s {
		_, ok := symbols[b]
		if ok {
			stack.Push(b)
		} else {
			if stack.Len() == 0 {
				return false
			}
			if b != symbols[stack.Peek().(rune)] {
				return false
			}
			_, _ = stack.Pop()
		}
	}

	return stack.Len() == 0
}
