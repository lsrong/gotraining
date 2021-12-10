package list

import (
	"errors"
	"fmt"
)

type Node struct {
	data interface{}
	next *Node

	list *List
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Data() interface{} {
	return n.data
}

type List struct {
	size int
	head *Node
	tail *Node
}

func New() *List {
	return new(List)
}

// InsertAfter 在指定的元素下一位插入一个新元素.
func (l *List) InsertAfter(data interface{}, element *Node) *Node {
	n := Node{}
	n.data = data
	if element == nil {
		// nil时从头部插入元素
		//    head											tail
		//	    V											 V
		//  [NEW].Next -> [Node2].Next -> [Node1].Next ->.[Node0].Next -> nil
		if l.size == 0 {
			l.tail = &n
		}
		n.next = l.head
		l.head = &n
	} else {
		// 插入到其中一个元素的下一位置
		//     head	          element						 tail
		//	    V				V							  V
		//  [Node2].Next -> [Element].Next -> [NEW].Next ->.[Node0].Next -> nil
		if element.next == nil {
			l.tail = &n
		}
		n.next = element.next
		element.next = &n

	}
	l.size++

	return &n
}

// RemoveAfter 移除指定元素的下一位元素.
func (l *List) RemoveAfter(element *Node) (*Node, error) {
	if l.Len() == 0 {
		return nil, errors.New("list is empty")
	}
	var delNode *Node
	if element == nil {
		// 从链表头部删除元素
		//    head(delete)					   tail
		//	    V								V
		//   [Node2].Next -> [Node1].Next ->.[Node0].Next -> nil
		if l.head.next == nil {
			// the last one is nil.
			l.tail = nil
		}
		delNode = l.head
		l.head = l.head.next

	} else {
		if element.next == nil {
			return nil, fmt.Errorf("can't removeAfter the last Node[%v]", element)
		}
		//    head					   		    delete		    tail
		//	    V								  V			     V
		//   [Node3].Next -> [Element].Next -> [Node1].Next -> [Node0].Next -> nil
		delNode = element.next
		element.next = element.next.next
		if element.next == nil {
			l.tail = element
		}
	}

	l.size--

	return delNode, nil
}

func (l *List) Len() int {
	return l.size
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Tail() *Node {
	return l.tail
}
