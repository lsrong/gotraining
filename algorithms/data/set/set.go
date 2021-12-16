package set

import (
	"errors"
	"github.com/learning_golang/algorithms/data/list"
)

var ErrNotFound = errors.New("can't find data in Set")

// Set 集合：相关成员的无序组合，且每个成员在集合中仅出现一次。
type Set struct {
	list list.List
}

// NewSet 生成新集合
func NewSet() *Set {
	return &Set{
		list: list.List{},
	}
}

// Insert 将一个成员插入集合中
func (s *Set) Insert(data interface{}) {
	if s.IsMember(data) {
		return
	}
	s.list.InsertAfter(data, nil)
}

// Remove 删除集合中的某个成员
func (s *Set) Remove(data interface{}) error {
	// 找到需要删除位置的前一个元素
	var prev *list.Node
	m := s.list.Head()
	for ; m != nil; m = m.Next() {
		if m.Data() == data {
			break
		}
		prev = m
	}
	if m == nil {
		return ErrNotFound
	}

	_, err := s.list.RemoveAfter(prev)
	if err != nil {
		return ErrNotFound
	}

	return nil
}

// IsMember 判断成员是否已经存在于集合中.
func (s *Set) IsMember(data interface{}) bool {
	m := s.list.Head()
	if m != nil {
		if m.Data() == data {
			return true
		}
		m = m.Next()
	}

	return false
}

func (s *Set) Size() int {
	return s.list.Len()
}

func (s *Set) IsSubset(subset *Set) bool {
	if subset.Size() > s.Size() {
		return false
	}

	for m := subset.list.Head(); m != nil; m = m.Next() {
		if !s.IsMember(m.Data()) {
			return false
		}
	}
	return true
}

func (s *Set) IsEqual(equal *Set) bool {
	if s.Size() != equal.Size() {
		return false
	}

	return s.IsSubset(equal)
}

// Union 合并两个集合，并集操作。
func Union(s1, s2 *Set) *Set {
	s := NewSet()
	// s1：所有成员添加到新集合中
	for m := s1.list.Head(); m != nil; m = m.Next() {
		s.list.InsertAfter(m.Data(), s.list.Tail())
	}

	// s2：先判断成员是否已在新集合中，不存在才添加
	for m := s2.list.Head(); m != nil; m = m.Next() {
		if s.IsMember(m.Data()) {
			continue
		}

		s.list.InsertAfter(m.Data(), s.list.Tail())
	}

	return s
}

func Intersection(s1, s2 *Set) *Set {
	s := NewSet()
	for m := s1.list.Head(); m != nil; m = m.Next() {
		if s2.IsMember(m.Data()) {
			s.list.InsertAfter(m.Data(), s.list.Tail())
		}
	}

	return s
}

func Difference(s1, s2 *Set) *Set {
	s := NewSet()
	for m := s1.list.Head(); m != nil; m = m.Next() {
		if !s2.IsMember(m.Data()) {
			s.list.InsertAfter(m.Data(), s.list.Tail())
		}
	}
	return s
}
