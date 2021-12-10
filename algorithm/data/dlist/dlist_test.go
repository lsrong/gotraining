package dlist

import "testing"

func checkListLen(t *testing.T, l *DList, len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("l.Len() Got %d, Expect %d", n, len)
		return false
	}

	return true
}

func checkList(t *testing.T, l *DList, es []*DNode) {
	root := &l.root
	if !checkListLen(t, l, len(es)) {
		return
	}
	if len(es) == 0 {
		if l.root.next != nil && l.root.next != root || l.root.prev != nil && l.root.prev != root {
			t.Errorf("l.root.prev = %p, l.root.next = %p; both should both be nil or %p", l.root.prev, l.root.next, root)
		}
		return
	}

	for i, e := range es {
		//next := root
		//Next := (*DNode)(nil)
		//if i < len(es)-1 {
		//	next = es[i+1]
		//	Next = next
		//}
		//if n := e.next; n != next {
		//	t.Errorf("elt[%d](%p).next = %p, Expect %p", i, e, n, next)
		//	return
		//}
		//if n := e.Next(); n != Next {
		//	t.Errorf("elt[%d](%p).Next() = %p, Expect %p", i, e, n, Next)
		//	return
		//}
		prev := root
		Prev := (*DNode)(nil)
		if i > 0 {
			prev = es[i-1]
			Prev = prev
		}
		if p := e.prev; p != prev {
			t.Errorf("elt[%d](%p).prev = %p, want %p", i, e, p, prev)
		}
		if p := e.Prev(); p != Prev {
			t.Errorf("elt[%d](%p).Prev() = %p, want %p", i, e, p, Prev)
		}

		next := root
		Next := (*DNode)(nil)
		if i < len(es)-1 {
			next = es[i+1]
			Next = next
		}
		if n := e.next; n != next {
			t.Errorf("elt[%d](%p).next = %p, want %p", i, e, n, next)
		}
		if n := e.Next(); n != Next {
			t.Errorf("elt[%d](%p).Next() = %p, want %p", i, e, n, Next)
		}
	}
}

func TestList(t *testing.T) {
	l := New()
	//l.InsertAfter("test", nil)
	checkList(t, l, []*DNode{})

	// Single
	e := l.PushFront(1)
	checkList(t, l, []*DNode{e})
	l.Remove(e)
	checkList(t, l, []*DNode{})

	// Multiple 4 3 1 2
	e1 := l.PushFront(1)
	e2 := l.InsertAfter(1, e1)
	e3 := l.PushFront(1)
	e4 := l.InsertBefore(1, e3)
	checkList(t, l, []*DNode{e4, e3, e1, e2})

	sum := 0
	for ptr := l.Front(); ptr != nil; ptr = ptr.Next() {
		if i, ok := ptr.Value().(int); ok {
			sum += i
			t.Log(i)
		}
	}
	if sum != 4 {
		t.Errorf("sum over l = %d, want 4", sum)
	}
}
