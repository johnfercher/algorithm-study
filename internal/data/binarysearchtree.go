package data

import "fmt"

type IntNode struct {
	Value    int
	Previous *IntNode
	Left     *IntNode
	Right    *IntNode
}

func (s *IntNode) IsLeaf() bool {
	if s == nil {
		return false
	}

	return s.Right == nil && s.Left == nil
}

func (s *IntNode) Print() {
	if s == nil {
		return
	}

	if s.IsLeaf() {
		fmt.Printf("(%d)\n", s.Value)
		return
	}

	if s.Right != nil {
		fmt.Printf("(%d) -> %d\n", s.Value, s.Right.Value)
		s.Right.Print()
	}

	if s.Left != nil {
		fmt.Printf("%d <- (%d)\n", s.Left.Value, s.Value)
		s.Left.Print()
	}

	fmt.Println()
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

func (s *BinaryIntSearchTree) Root() (int, bool) {
	if s.root == nil {
		return 0, false
	}

	return s.root.Value, true
}

func (s *BinaryIntSearchTree) Add(value int) bool {
	s.length++
	if s.root == nil {
		s.root = &IntNode{
			Value: value,
		}
		return true
	}

	closerNode, sameValue := s.getCloserNode(value)
	if sameValue {
		return false
	}

	newNode := &IntNode{
		Value:    value,
		Previous: closerNode,
	}

	if value > closerNode.Value {
		closerNode.Right = newNode
	} else {
		closerNode.Left = newNode
	}

	return true
}

func (s *BinaryIntSearchTree) Remove(value int) bool {
	if s.root == nil {
		return false
	}

	s.length--

	// Need to implement
	if s.root.Value == value {
		return false
	}

	if s.root.Value == value {
	}

	closerNode, sameValue := s.getCloserNode(value)
	if !sameValue {
		return false
	}

	s.removeNodeFromBranch(closerNode)
	return true
}

func (s *BinaryIntSearchTree) Exists(value int) bool {
	_, sameValue := s.getCloserNode(value)
	if !sameValue {
		return false
	}

	return true
}

func (s *BinaryIntSearchTree) Print() {
	s.root.Print()
}

func (s *BinaryIntSearchTree) removeNodeFromBranch(node *IntNode) {
	if node.Value > node.Previous.Value {
		node.Previous.Right = node.Right
		return
	}

	node.Previous.Left = node.Left
}

func (s *BinaryIntSearchTree) getCloserNode(value int) (node *IntNode, sameValue bool) {
	nodeIterator := s.root
	closerNode := s.root

	for nodeIterator != nil {
		if value == nodeIterator.Value {
			return nodeIterator, true
		}

		closerNode = nodeIterator

		if value > nodeIterator.Value {
			nodeIterator = nodeIterator.Right
		} else {
			nodeIterator = nodeIterator.Left
		}
	}

	return closerNode, false
}
