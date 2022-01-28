package data

import "algorithm-study/internal/hash"

type IntHashSet struct {
	maxLength      int
	length         int
	arrayValue     []int
	arrayMemoryMap []bool
}

func NewIntHashSet(size int) *IntHashSet {
	var arrValue []int
	var arrOccupation []bool
	for i := 0; i < size; i++ {
		arrValue = append(arrValue, 0)
		arrOccupation = append(arrOccupation, false)
	}

	return &IntHashSet{
		arrayValue:     arrValue,
		arrayMemoryMap: arrOccupation,
		maxLength:      size,
		length:         0,
	}
}

func (s *IntHashSet) Add(key string, value int) {
	hashKey := hash.Generate(key, s.maxLength)
	s.arrayValue[hashKey] = value
	s.arrayMemoryMap[hashKey] = true
	s.length++
}

func (s *IntHashSet) Remove(key string) {
	hashKey := hash.Generate(key, s.maxLength)
	s.arrayMemoryMap[hashKey] = false
	s.length--
}

func (s *IntHashSet) Get(key string) (int, bool) {
	hashKey := hash.Generate(key, s.maxLength)
	if !s.arrayMemoryMap[hashKey] {
		return 0, false
	}

	return s.arrayValue[hashKey], true
}

func (s *IntHashSet) Length() int {
	return s.length
}
