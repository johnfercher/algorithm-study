package data

import "fmt"

type StackIntArray struct {
	arr    []int
	length int
}

func NewStackIntArray() *StackIntArray {
	return &StackIntArray{}
}

func (s *StackIntArray) Length() int {
	return s.length
}

func (s *StackIntArray) Push(value int) {
	s.length++
	s.arr = append(s.arr, value)
}

func (s *StackIntArray) Pop() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	value := s.arr[s.length-1]
	s.arr = s.arr[:s.length-1]

	s.length--
	return value, true
}

func (s *StackIntArray) Peek() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	return s.arr[s.length-1], true
}

func (s *StackIntArray) Print() {
	fmt.Println(s.arr)
}
