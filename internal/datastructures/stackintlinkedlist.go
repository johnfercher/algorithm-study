package datastructures

import "fmt"

type StackNodeInt struct {
	value int
	next  *StackNodeInt
}

func (s *StackNodeInt) Print() {
	fmt.Printf("(%d) -> ", s.value)
	if s.next != nil {
		s.next.Print()
	} else {
		fmt.Println("NULL")
	}
}

func NewStackNodeInt(value int, next *StackNodeInt) *StackNodeInt {
	return &StackNodeInt{
		value: value,
		next:  next,
	}
}

type StackIntLinkedList struct {
	top    *StackNodeInt
	length int
}

func NewStackIntLinkedList() *StackIntLinkedList {
	return &StackIntLinkedList{
		top:    nil,
		length: 0,
	}
}

func (s *StackIntLinkedList) Length() int {
	return s.length
}

func (s *StackIntLinkedList) Push(value int) {
	s.length++
	s.top = NewStackNodeInt(value, s.top)
}

func (s *StackIntLinkedList) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}

	value := s.top.value
	s.length--
	s.top = s.top.next

	return value, true
}

func (s *StackIntLinkedList) Peek() (int, bool) {
	if s.top == nil {
		return 0, false
	}

	return s.top.value, true
}

func (s *StackIntLinkedList) Print() {
	if s.top != nil {
		s.top.Print()
	}
}
