package data

import "fmt"

type IntArray struct {
	length int
	data   []int
}

func NewIntArray() *IntArray {
	return &IntArray{}
}

func (s *IntArray) Len() int {
	return s.length
}

func (s *IntArray) Print() {
	fmt.Println(s.data)
}

func (s *IntArray) Push(a int) {
	s.length++
	s.data = append(s.data, a)
}

func (s *IntArray) PushFront(a int) {
	s.length++

	s.data = append(s.data, 0)
	for i := 0; i < len(s.data)-1; i++ {
		s.data[i+1] = s.data[i]
	}

	s.data[0] = a
}

func (s *IntArray) Pop() int {
	value := s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length--
	return value
}

func (s *IntArray) PopFront() int {
	value := s.data[0]
	s.data = s.data[1:]
	s.length--
	return value
}

func (s *IntArray) At(index int) int {
	return s.data[index]
}
