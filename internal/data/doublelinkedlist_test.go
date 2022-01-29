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

func TestIntDoubleLinkedList_PushBack(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	value := 0
	ok := false

	// Empty Data
	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	// Push Back 1
	sut.PushBack(1)

	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	assert.Equal(t, 1, sut.Length())

	// Push Back 2
	sut.PushBack(2)

	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	assert.Equal(t, 2, sut.Length())

	// Push Back 3
	sut.PushBack(3)

	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	assert.Equal(t, 3, sut.Length())
}

func TestIntDoubleLinkedList_PushFront(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	value := 0
	ok := false

	// Empty Data
	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	// Push Front 1
	sut.PushFront(1)

	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	assert.Equal(t, 1, sut.Length())

	// Push Front 2
	sut.PushFront(2)

	value, ok = sut.Head()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	assert.Equal(t, 2, sut.Length())

	// Push Front 3
	sut.PushFront(3)

	value, ok = sut.Head()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	assert.Equal(t, 3, sut.Length())
}

func TestIntDoubleLinkedList_PushAt(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	value := 0
	ok := false

	// Empty Data
	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	// Push At Begin
	sut.PushAt(1, 0)

	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	assert.Equal(t, 1, sut.Length())

	// Push At Begin
	sut.PushAt(2, 0)

	value, ok = sut.Head()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	assert.Equal(t, 2, sut.Length())

	// Push At End
	sut.PushAt(3, 1)

	value, ok = sut.Head()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	assert.Equal(t, 3, sut.Length())

	// Push At Middle
	sut.PushAt(4, 1)

	value, ok = sut.Head()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	assert.Equal(t, 4, sut.Length())
}

func TestIntDoubleLinkedList_PopBack(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	value := 0
	ok := false

	sut.PushFront(1)
	sut.PushFront(2)
	sut.PushFront(3)

	// First Pop
	value, ok = sut.PopBack()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Head()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	assert.Equal(t, 2, sut.Length())

	// Second Pop
	value, ok = sut.PopBack()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	value, ok = sut.Head()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	assert.Equal(t, 1, sut.Length())

	// Last Pop
	value, ok = sut.PopFront()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	assert.Equal(t, 0, sut.Length())

	// Invalid Pop
	value, ok = sut.PopFront()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	assert.Equal(t, 0, sut.Length())
}

func TestIntDoubleLinkedList_PopFront(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	value := 0
	ok := false

	sut.PushFront(1)
	sut.PushFront(2)
	sut.PushFront(3)

	// First Pop
	value, ok = sut.PopFront()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	value, ok = sut.Head()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	assert.Equal(t, 2, sut.Length())

	// Second Pop
	value, ok = sut.PopFront()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	assert.Equal(t, 1, sut.Length())

	// Last Pop
	value, ok = sut.PopFront()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	assert.Equal(t, 0, sut.Length())

	// Invalid Pop
	value, ok = sut.PopFront()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	assert.Equal(t, 0, sut.Length())
}

func TestIntDoubleLinkedList_PopAt(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	value := 0
	ok := false

	sut.PushBack(1)
	sut.PushBack(2)
	sut.PushBack(3)
	sut.PushBack(4)

	// Pop From Middle
	value, ok = sut.PopAt(2)
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 4, value)
	assert.True(t, ok)

	// Pop From End
	value, ok = sut.PopAt(2)
	assert.Equal(t, 4, value)
	assert.True(t, ok)

	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	assert.Equal(t, 2, sut.Length())

	// Pop From Begin
	value, ok = sut.PopAt(0)
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Head()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	assert.Equal(t, 1, sut.Length())

	// Last Pop
	value, ok = sut.PopAt(0)
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	assert.Equal(t, 0, sut.Length())

	// Invalid Pop
	value, ok = sut.PopAt(0)
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	assert.Equal(t, 0, sut.Length())
}

func TestIntDoubleLinkedList_At(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

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
	assert.Zero(t, value)
}
