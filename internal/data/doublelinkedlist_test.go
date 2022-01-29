package data_test

import (
	"algorithm-study/internal/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIntLinkedList(t *testing.T) {
	// Act
	sut := data.NewDoubleIntLinkedList()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, 0, sut.Length())
}

func TestIntDoubleLinkedList_PushFront(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	value := 0
	ok := false

	// Act & Assert
	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	sut.PushFront(1)
	assert.Equal(t, 1, sut.Length())
	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	sut.PushFront(2)
	assert.Equal(t, 2, sut.Length())
	value, ok = sut.Head()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	sut.PushFront(3)
	assert.Equal(t, 3, sut.Length())
	value, ok = sut.Head()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
}

func TestIntDoubleLinkedList_PushBack(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	value := 0
	ok := false

	// Act & Assert
	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	sut.PushBack(1)
	assert.Equal(t, 1, sut.Length())
	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	sut.PushBack(2)
	assert.Equal(t, 2, sut.Length())
	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	sut.PushBack(3)
	assert.Equal(t, 3, sut.Length())
	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
}

func TestIntDoubleLinkedList_PushAt(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	sut.PushBack(1)
	sut.PushBack(2)
	sut.PushBack(3)
	sut.PushBack(4)
	sut.PushBack(5)
	sut.PushBack(6)

	value := 0
	ok := false

	// Act & Assert
	// Add Forward
	sut.PushAt(10, 1)
	assert.Equal(t, 7, sut.Length())
	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 6, value)
	assert.True(t, ok)

	// Add Backward
	sut.PushAt(10, 5)
	assert.Equal(t, 8, sut.Length())
	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 6, value)
	assert.True(t, ok)
}

func TestIntDoubleLinkedList_PopFront(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	sut.PushBack(1)
	sut.PushBack(2)
	sut.PushBack(3)

	value, ok := sut.PopFront()

	// Act
	assert.Equal(t, 2, sut.Length())
	assert.Equal(t, 1, value)
	assert.True(t, ok)
}

func TestIntDoubleLinkedList_PopBack(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	sut.PushBack(1)
	sut.PushBack(2)
	sut.PushBack(3)

	value, ok := sut.PopBack()

	// Act
	assert.Equal(t, 2, sut.Length())
	assert.Equal(t, 3, value)
	assert.True(t, ok)
}

func TestIntDoubleLinkedList_PopAt(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	sut.PushBack(1)
	sut.PushBack(2)
	sut.PushBack(3)

	value, ok := sut.PopAt(1)

	// Act
	assert.Equal(t, 2, sut.Length())
	assert.Equal(t, 2, value)
	assert.True(t, ok)
}

func TestIntDoubleLinkedList_At(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	sut.PushBack(1)
	sut.PushBack(2)
	sut.PushBack(3)
	sut.PushBack(4)
	sut.PushBack(5)

	value, ok := sut.At(3)

	// Act
	assert.Equal(t, 5, sut.Length())
	assert.Equal(t, 4, value)
	assert.True(t, ok)
}
