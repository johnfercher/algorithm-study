package data_test

import (
	"algorithm-study/internal/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSingleIntLinkedList(t *testing.T) {
	// Act
	sut := data.NewSingleIntLinkedList()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, 0, sut.Length())
}

func TestSingleIntLinkedList_PushFront(t *testing.T) {
	// Arrange
	sut := data.NewSingleIntLinkedList()

	// Act
	sut.PushFront(0)
	sut.PushFront(1)
	sut.PushFront(2)

	// Assert
	assert.Equal(t, 3, sut.Length())
	value, _ := sut.PopFront()
	assert.Equal(t, 2, value)
	value, _ = sut.PopFront()
	assert.Equal(t, 1, value)
	value, _ = sut.PopFront()
	assert.Equal(t, 0, value)
}

func TestSingleIntLinkedList_PopFront(t *testing.T) {
	// Arrange
	sut := data.NewSingleIntLinkedList()

	sut.PushFront(0)
	sut.PushFront(1)
	sut.PushFront(2)

	// Act & Assert
	assert.Equal(t, 3, sut.Length())
	value, _ := sut.PopFront()
	assert.Equal(t, 2, value)
	value, _ = sut.PopFront()
	assert.Equal(t, 1, value)
	value, _ = sut.PopFront()
	assert.Equal(t, 0, value)
	_, ok := sut.PopFront()
	assert.False(t, ok)
}

func TestSingleIntLinkedList_Tail(t *testing.T) {
	// Arrange
	sut := data.NewSingleIntLinkedList()

	// Act & Assert
	value, ok := sut.Tail()
	assert.False(t, ok)
	assert.Zero(t, value)

	sut.PushFront(1)
	value, ok = sut.Tail()
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	sut.PushFront(2)
	value, ok = sut.Tail()
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	sut.PopFront()
	value, ok = sut.Tail()
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	sut.PopFront()
	value, ok = sut.Tail()
	assert.False(t, ok)
	assert.Zero(t, value)
}

func TestSingleIntLinkedList_Head(t *testing.T) {
	// Arrange
	sut := data.NewSingleIntLinkedList()

	// Act & Assert
	value, ok := sut.Head()
	assert.False(t, ok)
	assert.Zero(t, value)

	sut.PushFront(1)
	value, ok = sut.Head()
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	sut.PushFront(2)
	value, ok = sut.Head()
	assert.True(t, ok)
	assert.Equal(t, 2, value)

	sut.PopFront()
	value, ok = sut.Head()
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	sut.PopFront()
	value, ok = sut.Head()
	assert.False(t, ok)
	assert.Zero(t, value)
}

func TestSingleIntLinkedList_PushBack(t *testing.T) {
	// Arrange
	sut := data.NewSingleIntLinkedList()

	// Act & Assert
	value, ok := sut.Tail()
	assert.False(t, ok)
	assert.Zero(t, value)

	sut.PushBack(1)
	value, ok = sut.Tail()
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	sut.PushBack(2)
	value, ok = sut.Tail()
	assert.True(t, ok)
	assert.Equal(t, 2, value)

	sut.PushBack(3)
	value, ok = sut.Tail()
	assert.True(t, ok)
	assert.Equal(t, 3, value)
}

func TestSingleIntLinkedList_PushAt(t *testing.T) {
	// Arrange
	sut := data.NewSingleIntLinkedList()

	// Act & Assert
	value, ok := sut.Tail()
	assert.False(t, ok)
	assert.Zero(t, value)

	sut.PushAt(1, 0)
	value, ok = sut.Head()
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	sut.PushAt(2, 0)
	value, ok = sut.Head()
	assert.True(t, ok)
	assert.Equal(t, 2, value)

	sut.PushAt(3, 0)
	value, ok = sut.Head()
	assert.True(t, ok)
	assert.Equal(t, 3, value)

	sut.PushAt(4, 1)
	value, ok = sut.Head()
	assert.True(t, ok)
	assert.Equal(t, 3, value)
}

func TestSingleIntLinkedList_At(t *testing.T) {
	// Arrange
	sut := data.NewSingleIntLinkedList()

	sut.PushBack(1)
	sut.PushBack(2)
	sut.PushBack(3)

	// Act & Assert
	value, ok := sut.At(0)
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	value, ok = sut.At(1)
	assert.True(t, ok)
	assert.Equal(t, 2, value)

	value, ok = sut.At(2)
	assert.True(t, ok)
	assert.Equal(t, 3, value)

	value, ok = sut.At(3)
	assert.False(t, ok)
	assert.Equal(t, 0, value)
}

func TestSingleIntLinkedList_PopBack(t *testing.T) {
	// Arrange
	sut := data.NewSingleIntLinkedList()

	sut.PushFront(0)
	sut.PushFront(1)
	sut.PushFront(2)

	// Act & Assert
	assert.Equal(t, 3, sut.Length())
	value, _ := sut.PopBack()
	assert.Equal(t, 0, value)
	value, _ = sut.PopBack()
	assert.Equal(t, 1, value)
	value, _ = sut.PopBack()
	assert.Equal(t, 2, value)
	_, ok := sut.PopBack()
	assert.False(t, ok)
}

func TestSingleIntLinkedList_PopAtBack(t *testing.T) {
	// Arrange
	sut := data.NewSingleIntLinkedList()

	sut.PushFront(0)
	sut.PushFront(1)
	sut.PushFront(2)

	// Act & Assert
	assert.Equal(t, 3, sut.Length())
	value, _ := sut.PopAt(2)
	assert.Equal(t, 0, value)
	value, _ = sut.PopBack()
	assert.Equal(t, 1, value)
	value, _ = sut.PopBack()
	assert.Equal(t, 2, value)
	_, ok := sut.PopBack()
	assert.False(t, ok)
}
