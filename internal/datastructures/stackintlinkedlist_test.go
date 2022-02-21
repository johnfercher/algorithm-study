package datastructures_test

import (
	"algorithm-study/internal/datastructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStackIntLinkedList(t *testing.T) {
	// Act
	sut := datastructures.NewStackIntLinkedList()

	// Assert
	assert.NotNil(t, sut)
	assert.Zero(t, sut.Length())
	value, ok := sut.Peek()
	assert.Zero(t, value)
	assert.False(t, ok)
}

func TestStackIntLinkedList_Push(t *testing.T) {
	// Arrange
	sut := datastructures.NewStackIntLinkedList()

	value := 0
	ok := false

	// Push 1
	sut.Push(1)
	assert.Equal(t, 1, sut.Length())
	value, ok = sut.Peek()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	// Push 2
	sut.Push(2)
	assert.Equal(t, 2, sut.Length())
	value, ok = sut.Peek()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	// Push 3
	sut.Push(3)
	assert.Equal(t, 3, sut.Length())
	value, ok = sut.Peek()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
}

func TestStackIntLinkedList_Peek(t *testing.T) {
	// Arrange
	sut := datastructures.NewStackIntLinkedList()

	sut.Push(1)
	sut.Push(2)
	sut.Push(3)

	// Act
	value, ok := sut.Peek()

	// Assert
	assert.Equal(t, 3, value)
	assert.True(t, ok)
}

func TestStackIntLinkedList_Pop(t *testing.T) {
	// Arrange
	sut := datastructures.NewStackIntLinkedList()

	value := 0
	ok := false

	sut.Push(1)
	sut.Push(2)
	sut.Push(3)

	// Pop 1
	value, ok = sut.Pop()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
	assert.Equal(t, 2, sut.Length())

	// Pop 2
	value, ok = sut.Pop()
	assert.Equal(t, 2, value)
	assert.True(t, ok)
	assert.Equal(t, 1, sut.Length())

	// Pop 3
	value, ok = sut.Pop()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	assert.Equal(t, 0, sut.Length())

	// Invalid Pop
	value, ok = sut.Pop()
	assert.Zero(t, value)
	assert.False(t, ok)
	assert.Equal(t, 0, sut.Length())
}

func TestStackIntLinkedList_Length(t *testing.T) {
	// Arrange
	sut := datastructures.NewStackIntLinkedList()

	sut.Push(1)
	sut.Push(1)
	sut.Push(1)

	// Act & Assert
	assert.Equal(t, 3, sut.Length())
}
