package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	s := NewQueue()
	const items = 5
	var orgData string

	// Push
	i := 0
	first := fmt.Sprintf("Name %d", i)
	for ; i < items; i++ {
		name := fmt.Sprintf("Name %d", i)
		orgData += name
		s.Enqueue(name)
	}

	if s.Len() != items {
		t.Fatalf("queue.Enqueue(), Len Got %d, Expected %d", s.Len(), items)
	}

	// Peek
	head := s.Peek().(string)
	if head != first {
		t.Fatalf("queue.Peek() Got %s, Expectd %s", head, first)
	}

	// Pop
	var queueData string
	for i := 0; i < items; i++ {
		name, err := s.Dequeue()
		if err != nil {
			t.Fatalf("queue.Dequeue(). should not be err: %v", err)
		}

		queueData += name.(string)
	}

	if queueData != orgData {
		t.Fatalf("queue Got %s, Expected %s", queueData, orgData)
	}
	t.Logf("Queue is %s", queueData)
}
