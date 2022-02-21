package datastructures

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

func (s *BinaryIntSearchTree) BreadthFirstSearch() []int {
	var list []int
	var queue []*IntNode

	currentNode := s.root
	queue = append(queue, currentNode)

	for len(queue) > 0 {
		currentNode := queue[0]
		list = append(list, currentNode.Value)
		queue = queue[1:]

		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}

		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}

	return list
}

func (s *BinaryIntSearchTree) DepthFirstSearchInOrder() []int {
	list := s.transverseInOrder(s.root)
	return list
}

func (s *BinaryIntSearchTree) DepthFirstSearchPreOrder() []int {
	list := s.transversePreOrder(s.root)
	return list
}

func (s *BinaryIntSearchTree) DepthFirstSearchPostOrder() []int {
	list := s.transversePostOrder(s.root)
	return list
}

func (s *BinaryIntSearchTree) transverseInOrder(node *IntNode) []int {
	var list []int
	if node.Left != nil {
		innerLeftList := s.transverseInOrder(node.Left)
		list = append(list, innerLeftList...)
	}

	list = append(list, node.Value)

	if node.Right != nil {
		innerRightList := s.transverseInOrder(node.Right)
		list = append(list, innerRightList...)
	}

	return list
}

func (s *BinaryIntSearchTree) transversePreOrder(node *IntNode) []int {
	var list []int
	list = append(list, node.Value)

	if node.Left != nil {
		innerLeftList := s.transversePreOrder(node.Left)
		list = append(list, innerLeftList...)
	}

	if node.Right != nil {
		innerRightList := s.transversePreOrder(node.Right)
		list = append(list, innerRightList...)
	}

	return list
}

func (s *BinaryIntSearchTree) transversePostOrder(node *IntNode) []int {
	var list []int

	if node.Left != nil {
		innerLeftList := s.transversePostOrder(node.Left)
		list = append(list, innerLeftList...)
	}

	if node.Right != nil {
		innerRightList := s.transversePostOrder(node.Right)
		list = append(list, innerRightList...)
	}

	list = append(list, node.Value)

	return list
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
