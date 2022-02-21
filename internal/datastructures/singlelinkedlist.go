package datastructures

import "fmt"

type IntSingleNode struct {
	value int
	next  *IntSingleNode
}

func NewSingleIntNode(value int, next *IntSingleNode) *IntSingleNode {
	return &IntSingleNode{
		value: value,
		next:  next,
	}
}

func (s *IntSingleNode) Print() {
	fmt.Printf("(%d) -> ", s.value)
	if s.next != nil {
		s.next.Print()
	} else {
		fmt.Println("NULL")
	}
}

type IntSingleLinkedList struct {
	length int
	head   *IntSingleNode
	tail   *IntSingleNode
}

func NewSingleIntLinkedList() *IntSingleLinkedList {
	return &IntSingleLinkedList{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (s *IntSingleLinkedList) PushFront(value int) {
	s.length++
	headWasNil := s.head == nil

	s.head = NewSingleIntNode(value, s.head)

	if headWasNil {
		s.tail = s.head
	}
}

func (s *IntSingleLinkedList) PushBack(value int) {
	s.length++

	if s.tail == nil {
		s.tail = NewSingleIntNode(value, nil)
		s.head = s.tail
		return
	}

	s.tail.next = NewSingleIntNode(value, nil)
	s.tail = s.tail.next
}

func (s *IntSingleLinkedList) PushAt(value int, position int) {
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
	nodeToAdd := &IntSingleNode{
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

func (s *IntSingleLinkedList) PopFront() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	s.length--
	popedValue := s.head.value

	s.head = s.head.next

	if s.length == 1 {
		s.tail = s.head
	}

	return popedValue, true
}

func (s *IntSingleLinkedList) PopBack() (int, bool) {
	return s.PopAt(s.length - 1)
}

func (s *IntSingleLinkedList) PopAt(position int) (int, bool) {
	if s.length == 0 || position > s.length-1 {
		return 0, false
	}

	if s.length == 1 || position == 0 {
		return s.PopFront()
	}

	node, ok := s.transverseUntil(position - 1)
	if !ok {
		return 0, false
	}

	if position == s.length-1 {
		s.tail = node
	}

	s.length--
	value := node.next.value
	node.next = node.next.next

	return value, true
}

func (s *IntSingleLinkedList) At(position int) (int, bool) {
	node, ok := s.transverseUntil(position)
	if !ok {
		return 0, false
	}

	return node.value, true
}

func (s *IntSingleLinkedList) Reverse() {
	if s.head == nil {
		return
	}

	var arrAux []*IntSingleNode

	node := s.head
	for node != nil {
		arrAux = append(arrAux, node)
		node = node.next
	}

	s.head = arrAux[len(arrAux)-1]
	s.tail = arrAux[0]
	s.tail.next = nil

	node = s.head
	for i := len(arrAux) - 2; i >= 0; i-- {
		nextNode := arrAux[i]
		node.next = nextNode
		node = nextNode
	}
}

func (s *IntSingleLinkedList) Length() int {
	return s.length
}

func (s *IntSingleLinkedList) Tail() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	return s.tail.value, true
}

func (s *IntSingleLinkedList) Head() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	return s.head.value, true
}

func (s *IntSingleLinkedList) Print() {
	if s.head != nil {
		s.head.Print()
	}
}

func (s *IntSingleLinkedList) transverseUntil(position int) (*IntSingleNode, bool) {
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
