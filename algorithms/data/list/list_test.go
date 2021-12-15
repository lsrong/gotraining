package list

import "testing"

func checkListLen(t *testing.T, l *List, len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("l.Len() Got %d, Expect %d", n, len)
		return false
	}

	return true
}

func checkList(t *testing.T, l *List, es []*Node) {
	if !checkListLen(t, l, len(es)) {
		return
	}
	if len(es) == 0 {
		if l.head != nil || l.tail != nil {
			t.Errorf("l.head = %p, l.tail = %p; both should both be nil", l.head, l.tail)
		}
		return
	}

	// [N3] ?-> [N2] ?-> [N1] ?-> [N0] -> nil

	for i, e := range es {
		next := (*Node)(nil)
		Next := (*Node)(nil)
		if i < len(es)-1 {
			next = es[i+1]
			Next = next
		}

		if n := e.next; n != next {
			t.Errorf("elt[%d](%p).next = %p, Expect %p", i, e, n, next)
			return
		}
		if n := e.Next(); n != Next {
			t.Errorf("elt[%d](%p).next = %p, Expect %p", i, e, n, Next)
			return
		}
	}
}

func TestList(t *testing.T) {
	l := New()
	//l.InsertAfter("test", nil)
	checkList(t, l, []*Node{})

	// Single
	e := l.InsertAfter(1, nil)
	checkList(t, l, []*Node{e})
	_, err := l.RemoveAfter(nil)
	if err != nil {
		t.Errorf("should be able RemoveAfter(nil) on the list: %v", err)
	}
	checkList(t, l, []*Node{})

	// Multiple
	e1 := l.InsertAfter(1, nil)
	e2 := l.InsertAfter(2, e1)
	e3 := l.InsertAfter(3, e2)
	e4 := l.InsertAfter(4, e3)
	//checkList(t, l, []*Node{e1, e3, e2, e4})
	checkList(t, l, []*Node{e1, e2, e3, e4})

	// remove e3
	_, err = l.RemoveAfter(e2)
	if err != nil {
		t.Errorf("should be able RemoveAfter(element *Node) on the list: %v", err)
	}
	checkList(t, l, []*Node{e1, e2, e4})

	// insert after e1
	e5 := l.InsertAfter(5, e1)
	checkList(t, l, []*Node{e1, e5, e2, e4})
}
