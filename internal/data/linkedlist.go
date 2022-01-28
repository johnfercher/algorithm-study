package data

import "fmt"

type IntNode struct {
	value int
	next  *IntNode
}

func NewIntNode(value int, next *IntNode) *IntNode {
	return &IntNode{
		value: value,
		next:  next,
	}
}

func (s *IntNode) Print() {
	fmt.Printf("(%d) -> ", s.value)
	if s.next != nil {
		s.next.Print()
	} else {
		fmt.Println("NULL")
	}
}

type IntLinkedList struct {
	length int
	head   *IntNode
	tail   *IntNode
}

func NewIntLinkedList() *IntLinkedList {
	return &IntLinkedList{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (s *IntLinkedList) PushFront(value int) {
	s.length++

	if s.head == nil {
		s.head = NewIntNode(value, nil)
		s.tail = s.head
		return
	}

	oldHead := *s.head
	s.head = NewIntNode(value, &oldHead)
}

func (s *IntLinkedList) PushBack(value int) {
	s.length++

	if s.tail == nil {
		s.tail = NewIntNode(value, nil)
		s.head = s.tail
		return
	}

	newTail := NewIntNode(value, nil)
	s.tail.next = newTail
	s.tail = newTail
}

func (s *IntLinkedList) PushAt(value int, position int) {
	if position > s.length {
		return
	}

	if position == 0 {
		s.PushFront(value)
		return
	}

	if position == s.length {
		s.PushBack(value)
		return
	}

	s.length++
	nodeToAdd := &IntNode{
		value: value,
	}

	node, ok := s.transverseUntil(position)
	if !ok {
		return
	}

	nodeToAdd.next = node.next
	node.next = nodeToAdd
	return
}

func (s *IntLinkedList) PopFront() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	s.length--
	popedValue := s.head.value

	if s.length != 0 {
		newHead := *s.head.next
		s.head = &newHead

		if s.length == 1 {
			s.tail = s.head
		}

		return popedValue, true
	}

	s.head = nil
	s.tail = nil
	return popedValue, true
}

func (s *IntLinkedList) PopBack() (int, bool) {
	return s.PopAt(s.length - 1)
}

func (s *IntLinkedList) PopAt(position int) (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	if s.length == 1 {
		return s.PopFront()
	}

	node, ok := s.transverseUntil(position - 1)
	if !ok {
		return 0, false
	}

	s.length--
	value := node.next.value
	node.next = nil

	return value, true
}

func (s *IntLinkedList) At(position int) (int, bool) {
	node, ok := s.transverseUntil(position)
	if !ok {
		return 0, false
	}

	return node.value, true
}

func (s *IntLinkedList) Length() int {
	return s.length
}

func (s *IntLinkedList) Tail() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	return s.tail.value, true
}

func (s *IntLinkedList) Head() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	return s.head.value, true
}

func (s *IntLinkedList) Print() {
	if s.head != nil {
		s.head.Print()
	}
}

func (s *IntLinkedList) transverseUntil(position int) (*IntNode, bool) {
	i := 0
	iteratorNode := s.head

	for i < s.length {
		if iteratorNode == nil {
			return nil, false
		}

		if i == position {
			return iteratorNode, true
		}

		iteratorNode = iteratorNode.next
		i++
	}

	return nil, false
}
