package dlist

// DNode 节点: 数据,上一个节点指针,下一个节点指针, 所属链表指针
type DNode struct {
	data interface{}
	next *DNode
	prev *DNode

	list *DList
}

// Value 返回节点数据
func (n *DNode) Value() interface{} {
	return n.data
}

// Next 返回节点的下一个指针节点
func (n *DNode) Next() *DNode {
	next := n.next
	if n.list != nil && next != &n.list.root {
		return next
	}
	return nil
}

// Prev 返回节点的上一个指针节点
func (n *DNode) Prev() *DNode {
	prev := n.prev
	if n.list != nil && prev != &n.list.root {
		return prev
	}
	return nil
}

// DList 链表, 包括哨兵节点和大小属性
type DList struct {
	root DNode // root 哨兵节点  root.next 为链表头部, root.prev 为链表尾部,
	size int   // 链表节点个数
}

// New 返回新的已初始化的链表指针.
func New() *DList {
	return new(DList).Init()
}

// Init 初始化链表.
func (l *DList) Init() *DList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.size = 0

	return l
}

// reInit 如果链表为空,重新初始化哨兵节点.
func (l *DList) reInit() {
	// 节点为空,需要重新设置哨兵节点.
	if l.root.next == nil {
		l.Init()
	}
}

// insertValue 生成新节点,统一添加在指定节点的下一位置.
func (l *DList) insertValue(v interface{}, at *DNode) *DNode {
	e := DNode{}
	e.data = v

	return l.insert(&e, at)
}

// insert 链表中插入新节点.
func (l *DList) insert(e, at *DNode) *DNode {
	//                  at         e
	//                   V         V
	// root <-> Prev.[Node0].Next <-> Prev.[Node2].Next <->Prev.[Node3].Next <-> root
	// 第一步 设置新节点的上下指针
	//                  at         		   e                  at.next
	//                   V                 V                    V
	// root <-> Prev.[Node0].Next <- Prev.[New].Next -> Prev.[Node2].Next <->Prev.[Node3].Next <-> root
	e.prev = at
	e.next = at.next
	// 第二步 设置e的上指针at的下指针为e, e的下指针at.next的上指针为e
	//                  at         		     e                  at.next
	//                   V                   V                    V
	// root <-> Prev.[Node0].Next <-> Prev.[New].Next <-> Prev.[Node2].Next <->Prev.[Node3].Next <-> root
	e.next.prev = e
	e.prev.next = e

	// 设置节点的所属链表
	e.list = l
	l.size++

	return e
}

// PushFront 向链表的头部添加新节点.
func (l *DList) PushFront(v interface{}) *DNode {
	l.reInit()
	return l.insertValue(v, &l.root)
}

// PushBack 向链表的尾部添加新节点.
func (l *DList) PushBack(v interface{}) *DNode {
	l.reInit()
	return l.insertValue(v, l.root.prev)
}

// InsertAfter 向链表的指定节点后面添加新节点.
func (l *DList) InsertAfter(v interface{}, element *DNode) *DNode {
	if element.list != l {
		return nil
	}
	return l.insertValue(v, element)
}

// InsertBefore 向链表的指定节点之前添加新节点.
func (l *DList) InsertBefore(v interface{}, element *DNode) *DNode {
	if element.list != l {
		return nil
	}
	return l.insertValue(v, element.prev)
}

// Remove 移除链表中的某个节点
func (l *DList) Remove(element *DNode) interface{} {
	if element.list != l {
		return nil
	}

	//                  e.prev         	    e                    e.next
	//                   V                  V                      V
	// root <-> Prev.[Node0].Next <-> Prev.[Node1].Next <-> Prev.[Node2].Next <->Prev.[Node3].Next <-> root
	// 第一步 e的上一个节点的下指针指向e.next
	//                  e.prev         	   e.next
	//                   V                  V
	// root <-> Prev.[Node0].Next -> Prev.[Node2].Next <->Prev.[Node3].Next <-> root
	element.prev.next = element.next

	// 第二步 e的下一个节点的上指针指向e.prev
	//                  e.prev         	   e.next
	//                   V                  V
	// root <-> Prev.[Node0].Next <-> Prev.[Node2].Next <->Prev.[Node3].Next <-> root
	element.next.prev = element.prev

	// 释放删除节点的上下指针, 所属链表指指针
	element.prev = nil
	element.next = nil
	element.list = nil
	l.size--

	return element.Value()
}

// Len 返回链表长度
func (l *DList) Len() int {
	return l.size
}

// Front 返回链表中的第一个节点指针
func (l *DList) Front() *DNode {
	if l.size == 0 {
		return nil
	}
	return l.root.next
}

// Back 返回链表中的最后一个节点指针
func (l *DList) Back() *DNode {
	if l.size == 0 {
		return nil
	}
	return l.root.prev
}
