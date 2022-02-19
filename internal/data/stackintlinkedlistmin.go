package data

type StackIntLinkedListMin struct {
	StackIntLinkedList
	min *StackNodeInt
}

func NewStackIntLinkedListMin() *StackIntLinkedListMin {
	return &StackIntLinkedListMin{}
}

func (s *StackIntLinkedListMin) Length() int {
	return s.StackIntLinkedList.Length()
}

func (s *StackIntLinkedListMin) Push(value int) {
	s.StackIntLinkedList.Push(value)
	if s.min == nil {
		s.min = NewStackNodeInt(value, nil)
	}

	if value < s.min.value {
		s.min = NewStackNodeInt(value, s.min)
	}
}

func (s *StackIntLinkedListMin) Pop() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	if s.length == 1 {
		s.min = nil
	}

	s.min = s.min.next

	return s.StackIntLinkedList.Pop()
}

func (s *StackIntLinkedListMin) Peek() (int, bool) {
	return s.StackIntLinkedList.Peek()
}

func (s *StackIntLinkedListMin) Print() {
	s.StackIntLinkedList.Print()
}

func (s *StackIntLinkedListMin) Min() (int, bool) {
	if s.Length() == 0 {
		return 0, false
	}

	return s.min.value, true
}
