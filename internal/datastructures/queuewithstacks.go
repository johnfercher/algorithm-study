package datastructures

type QueueWithStack struct {
	stack1, stack2 *StackIntLinkedList
	length         int
}

func NewQueueWithStack() *QueueWithStack {
	return &QueueWithStack{
		stack1: NewStackIntLinkedList(),
		stack2: NewStackIntLinkedList(),
	}
}

func (s *QueueWithStack) Length() int {
	return s.length
}

func (s *QueueWithStack) Enqueue(value int) {
	s.length++
	for {
		valueInStack, ok := s.stack1.Pop()
		if !ok {
			break
		}

		s.stack2.Push(valueInStack)
	}

	s.stack1.Push(value)

	for {
		valueInStack, ok := s.stack2.Pop()
		if !ok {
			break
		}

		s.stack1.Push(valueInStack)
	}
}

func (s *QueueWithStack) Dequeue() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	s.length--
	return s.stack1.Pop()
}

func (s *QueueWithStack) First() (int, bool) {
	return s.stack1.Peek()
}

func (s *QueueWithStack) Print() {
	s.stack1.Print()
}
