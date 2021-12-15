package list

import (
	"fmt"
	"testing"
)

func TestMergeListNode(t *testing.T) {
	val1 := []int{1, 2, 4}
	val2 := []int{1, 3, 4}
	resultVal := []int{1, 1, 2, 3, 4, 4}
	l1 := NewListNode(val1)
	l2 := NewListNode(val2)

	l3 := MergeListNode(l1, l2)
	var l3Val []int
	length := 0
	lenFunc := func(n *ListNode) error {
		l3Val = append(l3Val, n.Val)
		length++
		return nil
	}
	_ = Operate(l3, lenFunc)
	if length != len(resultVal) {
		t.Errorf("MergeListNode result is [%v] expect[%v]", l3Val, resultVal)
		return
	}
	i := 0
	checkFunc := func(n *ListNode) error {
		if n.Val != resultVal[i] {
			return fmt.Errorf("reaultVal[%d] is %d, but ListNode.Val is %d", i, resultVal[i], n.Val)
		}
		i++
		return nil
	}
	err := Operate(l3, checkFunc)
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveNthFromFront(t *testing.T) {
	val := []int{1, 2, 3, 4, 5}
	list := NewListNode(val)

	result := []int{1, 2, 4, 5}
	rmList := RemoveNthFromFront(list, 3)
	rmLen := Len(rmList)
	if rmLen != len(result) {
		t.Errorf("rmList len is %d expect %d", rmLen, len(result))
		return
	}

	index := 0
	checkFunc := func(n *ListNode) error {
		if n.Val != result[index] {
			return fmt.Errorf("rmList index[%d] is %d, expect %d", index, n.Val, result[index])
		}
		index++

		return nil
	}
	err := Operate(rmList, checkFunc)
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	val := []int{1, 2, 3, 4, 5}
	list := NewListNode(val)

	result := []int{1, 2, 3, 5}
	rmList := RemoveNthFromEnd(list, 2)
	rmLen := Len(rmList)
	if rmLen != len(result) {
		t.Errorf("rmList len is %d expect %d", rmLen, len(result))
		return
	}

	index := 0
	checkFunc := func(n *ListNode) error {
		if n.Val != result[index] {
			return fmt.Errorf("rmList index[%d] is %d, expect %d", index, n.Val, result[index])
		}
		index++

		return nil
	}
	err := Operate(rmList, checkFunc)
	if err != nil {
		t.Error(err)
	}
}
