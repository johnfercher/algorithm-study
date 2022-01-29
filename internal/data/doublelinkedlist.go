package data

import "fmt"

type IntDoubleNode struct {
	value  int
	next   *IntDoubleNode
	before *IntDoubleNode
}

func NewDoubleIntNode(value int, next *IntDoubleNode, before *IntDoubleNode) *IntDoubleNode {
	return &IntDoubleNode{
		value:  value,
		next:   next,
		before: before,
	}
}

func (s *IntDoubleNode) PrintForward() {
	fmt.Printf("(%d) -> ", s.value)
	if s.next != nil {
		s.next.PrintForward()
	} else {
		fmt.Println("NULL")
	}
}

func (s *IntDoubleNode) PrintBackward() {
	fmt.Printf("(%d) -> ", s.value)
	if s.before != nil {
		s.before.PrintBackward()
	} else {
		fmt.Println("NULL")
	}
}

type IntDoubleLinkedList struct {
	length int
	head   *IntDoubleNode
	tail   *IntDoubleNode
}

func NewDoubleIntLinkedList() *IntDoubleLinkedList {
	return &IntDoubleLinkedList{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (s *IntDoubleLinkedList) PushFront(value int) {
	s.length++

	if s.head == nil {
		s.head = NewDoubleIntNode(value, nil, nil)
		s.tail = s.head
		return
	}

	oldHead := s.head
	newHead := NewDoubleIntNode(value, oldHead, nil)
	oldHead.before = newHead

	s.head = newHead
}

func (s *IntDoubleLinkedList) PushBack(value int) {
	s.length++

	if s.tail == nil {
		s.tail = NewDoubleIntNode(value, nil, nil)
		s.head = s.tail
		return
	}

	newTail := NewDoubleIntNode(value, nil, s.tail)
	s.tail.next = newTail
	s.tail = newTail
}

func (s *IntDoubleLinkedList) PushAt(value int, position int) {
	if position > s.length {
		return
	}

	if position == 0 {
		s.PushFront(value)
		return
	}

	if position == s.length-1 {
		s.PushBack(value)
		return
	}

	s.length++
	nodeToAdd := &IntDoubleNode{
		value: value,
	}

	node, ok := s.transverseUntil(position)
	if !ok {
		return
	}

	nodeToAdd.next = node
	nodeToAdd.before = node.before

	node.before.next = nodeToAdd
	node.before = nodeToAdd

	return
}

func (s *IntDoubleLinkedList) PopFront() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	s.length--
	popedValue := s.head.value

	newHead := s.head.next
	newHead.before = nil
	s.head = newHead

	if s.length == 1 {
		s.tail = s.head
	}

	return popedValue, true
}

func (s *IntDoubleLinkedList) PopBack() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	s.length--
	popedValue := s.tail.value

	newTail := s.tail.before
	newTail.next = nil
	s.tail = newTail

	if s.length == 1 {
		s.tail = s.head
	}

	return popedValue, true
}

func (s *IntDoubleLinkedList) PopAt(position int) (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	if s.length == 1 {
		return s.PopFront()
	}

	if position == 0 {
		return s.PopFront()
	}

	if position == s.length-1 {
		return s.PopBack()
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

func (s *IntDoubleLinkedList) At(position int) (int, bool) {
	node, ok := s.transverseUntil(position)
	if !ok {
		return 0, false
	}

	return node.value, true
}

func (s *IntDoubleLinkedList) Length() int {
	return s.length
}

func (s *IntDoubleLinkedList) Tail() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	return s.tail.value, true
}

func (s *IntDoubleLinkedList) Head() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	return s.head.value, true
}

func (s *IntDoubleLinkedList) PrintForward() {
	fmt.Printf("Forward: ")
	if s.head != nil {
		s.head.PrintForward()
	}
}

func (s *IntDoubleLinkedList) PrintBackward() {
	fmt.Printf("Backward: ")
	if s.tail != nil {
		s.tail.PrintBackward()
	}
}

func (s *IntDoubleLinkedList) transverseUntil(position int) (*IntDoubleNode, bool) {
	if s.length <= 3 {
		return s.forwardTransverseUntil(position)
	}

	if position <= s.length/2 {
		return s.forwardTransverseUntil(position)
	}

	return s.backwardTransverseUntil(position)
}

func (s *IntDoubleLinkedList) forwardTransverseUntil(position int) (*IntDoubleNode, bool) {
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

func (s *IntDoubleLinkedList) backwardTransverseUntil(position int) (*IntDoubleNode, bool) {
	inversePosition := s.length - 1 - position
	i := 0
	iteratorNode := s.tail

	for i < s.length {
		if iteratorNode == nil {
			return nil, false
		}

		if i == inversePosition {
			return iteratorNode, true
		}

		iteratorNode = iteratorNode.before
		i++
	}

	return nil, false
}
