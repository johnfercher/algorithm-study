package datastructures_test

import (
	"algorithm-study/internal/datastructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBinaryIntSearchTree(t *testing.T) {
	// Act
	sut := datastructures.NewBinaryIntSearchTree()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, 0, sut.Length())
}

func TestBinaryIntSearchTree_Add(t *testing.T) {
	// Arrange
	sut := datastructures.NewBinaryIntSearchTree()

	value := 0
	ok := true

	// Zero Adds
	value, ok = sut.Root()
	assert.Zero(t, value)
	assert.False(t, ok)
	assert.Equal(t, 0, sut.Length())

	// First Add
	sut.Add(100)
	value, ok = sut.Root()
	assert.Equal(t, 100, value)
	assert.True(t, ok)
	assert.Equal(t, 1, sut.Length())

	// Second Add
	sut.Add(200)
	value, ok = sut.Root()
	assert.Equal(t, 100, value)
	assert.True(t, ok)
	assert.Equal(t, 2, sut.Length())
}

func TestBinaryIntSearchTree_Exists(t *testing.T) {
	// Arrange
	sut := datastructures.NewBinaryIntSearchTree()

	sut.Add(500)
	sut.Add(250)
	sut.Add(1000)

	exists := false

	// Doesnt Exists
	exists = sut.Exists(666)
	assert.False(t, exists)

	// Exists Root
	exists = sut.Exists(500)
	assert.True(t, exists)

	// Exists Left
	exists = sut.Exists(250)
	assert.True(t, exists)

	// Exists Right
	exists = sut.Exists(1000)
	assert.True(t, exists)
}

func TestBinaryIntSearchTree_BreadthFirstSearch(t *testing.T) {
	// Arrange
	sut := datastructures.NewBinaryIntSearchTree()

	sut.Add(500)
	sut.Add(250)
	sut.Add(1000)

	// Act
	nodes := sut.BreadthFirstSearch()

	// Assert
	assert.Equal(t, []int{500, 250, 1000}, nodes)
}

func TestBinaryIntSearchTree_DepthFirstSearchInOrder(t *testing.T) {
	// Arrange
	sut := datastructures.NewBinaryIntSearchTree()

	sut.Add(500)
	sut.Add(250)
	sut.Add(1000)

	// Act
	nodes := sut.DepthFirstSearchInOrder()

	// Assert
	assert.Equal(t, []int{250, 500, 1000}, nodes)
}

func TestBinaryIntSearchTree_DepthFirstSearchPreOrder(t *testing.T) {
	// Arrange
	sut := datastructures.NewBinaryIntSearchTree()

	sut.Add(500)
	sut.Add(250)
	sut.Add(1000)

	// Act
	nodes := sut.DepthFirstSearchPreOrder()

	// Assert
	assert.Equal(t, []int{500, 250, 1000}, nodes)
}

func TestBinaryIntSearchTree_DepthFirstSearchPostOrder(t *testing.T) {
	// Arrange
	sut := datastructures.NewBinaryIntSearchTree()

	sut.Add(500)
	sut.Add(250)
	sut.Add(1000)

	// Act
	nodes := sut.DepthFirstSearchPostOrder()

	// Assert
	assert.Equal(t, []int{250, 1000, 500}, nodes)
}
