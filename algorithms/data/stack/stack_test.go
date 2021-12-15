package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	const items = 5
	var orgData string
	var last string

	// Push
	for i := 0; i < items; i++ {
		name := fmt.Sprintf("Name %d", i)
		last = name
		orgData = name + orgData
		s.Push(name)
	}

	if s.Len() != items {
		t.Fatalf("stack.Push, Len Got %d, Expected %d", s.Len(), items)
	}

	// Peek
	head := s.Peek().(string)
	if head != last {
		t.Fatalf("stack.Peek Got %s, Expectd %s", head, last)
	}

	// Pop
	var testData string
	for i := 0; i < items; i++ {
		name, err := s.Pop()
		if err != nil {
			t.Fatalf("stack.Pop should not be err: %v", err)
		}

		testData += name.(string)
	}

	if testData != orgData {
		t.Fatalf("stack Got %s, Expected %s", testData, orgData)
	}

}
