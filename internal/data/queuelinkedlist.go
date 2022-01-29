package data

import "fmt"

type QueueNodeInt struct {
	value    int
	next     *QueueNodeInt
	previous *QueueNodeInt
}

func (s *QueueNodeInt) Print() {
	fmt.Printf("(%d) -> ", s.value)
	if s.next != nil {
		s.next.Print()
	} else {
		fmt.Println("NULL")
	}
}

type QueueIntLinkedList struct {
	last   *QueueNodeInt
	first  *QueueNodeInt
	length int
}

func NewQueueIntLinkedList() *QueueIntLinkedList {
	return &QueueIntLinkedList{
		last:   nil,
		length: 0,
	}
}

func (s *QueueIntLinkedList) Length() int {
	return s.length
}

func (s *QueueIntLinkedList) Enqueue(value int) {
	s.length++
	if s.last == nil {
		s.last = &QueueNodeInt{
			value: value,
		}

		s.first = s.last
		return
	}

	newLast := &QueueNodeInt{
		value:    value,
		previous: s.last,
		next:     nil,
	}

	s.last.next = newLast
	s.last = newLast
}

func (s *QueueIntLinkedList) Dequeue() (int, bool) {
	if s.first == nil {
		return 0, false
	}

	s.length--
	value := s.first.value

	if s.length == 0 {
		s.first = nil
		s.last = nil
		return value, true
	}

	s.first = s.first.next
	return value, true
}

func (s *QueueIntLinkedList) First() (int, bool) {
	if s.first == nil {
		return 0, false
	}

	return s.first.value, true
}

func (s *QueueIntLinkedList) Last() (int, bool) {
	if s.last == nil {
		return 0, false
	}

	return s.last.value, true
}

func (s *QueueIntLinkedList) Print() {
	if s.first != nil {
		s.first.Print()
	}
}
