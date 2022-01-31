package data

import "fmt"

type IntNode struct {
	Value int
	Left  *IntNode
	Right *IntNode
}

func (s *IntNode) Print() {
	fmt.Printf("Value(%d)\n", s.Value)
	fmt.Printf("\\/\n")
	if s.Right != nil {
		fmt.Printf("-> %d\n", s.Right.Value)
		s.Right.Print()
	}

	if s.Left != nil {
		fmt.Printf("%d <-\n", s.Left.Value)
		s.Left.Print()
	}
}

type BinaryIntSearchTree struct {
	root   *IntNode
	length int
}

func NewBinaryIntSearchTree() *BinaryIntSearchTree {
	return &BinaryIntSearchTree{}
}

func (s *BinaryIntSearchTree) Length() int {
	return s.length
}

func (s *BinaryIntSearchTree) Add(value int) bool {
	s.length++
	if s.root == nil {
		s.root = &IntNode{
			Value: value,
		}
		return true
	}

	previous, current, validInsertion := s.GetEdgeToAddValue(value)
	if !validInsertion {
		return false
	}

	newNode := &IntNode{
		Value: value,
	}

	if current != nil {
		if current.Value > value {
			newNode.Right = current
		} else {
			newNode.Left = current
		}
	}

	if value > previous.Value {
		previous.Right = newNode
	} else {
		previous.Left = newNode
	}

	return true
}

func (s *BinaryIntSearchTree) GetEdgeToAddValue(value int) (previous *IntNode, current *IntNode, validInsertion bool) {
	var oldParent *IntNode
	parent := s.root

	for parent != nil {
		if value == parent.Value {
			return nil, nil, false
		}

		oldParent = parent
		if value > parent.Value {
			parent = parent.Right
		} else {
			parent = parent.Left
		}
	}

	return oldParent, parent, true
}

func (s *BinaryIntSearchTree) Print() {
	s.root.Print()
}
