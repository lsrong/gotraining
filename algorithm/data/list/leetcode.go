package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode 根据值切片生成对应的链表
func NewListNode(val []int) *ListNode {
	head := &ListNode{}
	tail := head
	for _, v := range val {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}

	return head.Next
}

// MergeListNode 合并两个有序链表, 例如： l1: 1->2-4->nil, l2:1->3->4->nil, 返回list: 1->1->2->3->4->4->nil
func MergeListNode(l1, l2 *ListNode) *ListNode {
	// 利用哨兵的思想。
	prevHead := &ListNode{}
	tail := prevHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return prevHead.Next
}

// RemoveNthFromFront 删除链表的第N个结点.
func RemoveNthFromFront(list *ListNode, n int) *ListNode {
	if n <= 0 {
		return list
	}
	prev := &ListNode{0, list}
	tail := prev
	for i := 0; i < n-1; i++ {
		if tail.Next.Next == nil {
			return prev.Next
		}
		tail = tail.Next
	}
	tail.Next = tail.Next.Next

	return prev.Next
}

// RemoveNthFromEnd 删除链表的倒数第N个结点.
func RemoveNthFromEnd(list *ListNode, n int) *ListNode {
	if n <= 0 {
		return list
	}
	length := Len(list)
	if n > length {
		return list
	}
	head := &ListNode{0, list}
	tail := head
	for i := 0; i < length-n; i++ {
		tail = tail.Next
	}
	tail.Next = tail.Next.Next

	return head.Next
}

func Operate(list *ListNode, f func(n *ListNode) error) error {
	for ptr := list; ptr != nil; ptr = ptr.Next {
		if err := f(ptr); err != nil {
			return err
		}
	}

	return nil
}

func Len(list *ListNode) int {
	length := 0
	lenF := func(n *ListNode) error {
		length++
		return nil
	}
	_ = Operate(list, lenF)

	return length
}
