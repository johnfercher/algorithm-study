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

	elements := []int{}
	value := 0
	ok := false

	// Empty Data
	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	assert.Equal(t, 0, sut.Length())
	elements, ok = sut.GetThToLast(0)
	assert.Nil(t, elements)
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
	elements, ok = sut.GetThToLast(0)
	assert.Equal(t, []int{1}, elements)
	assert.True(t, ok)

	// Push Back 2
	sut.PushBack(2)

	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	assert.Equal(t, 2, sut.Length())
	elements, ok = sut.GetThToLast(0)
	assert.Equal(t, []int{1, 2}, elements)
	assert.True(t, ok)

	// Push Back 3
	sut.PushBack(3)

	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	assert.Equal(t, 3, sut.Length())
	elements, ok = sut.GetThToLast(0)
	assert.Equal(t, []int{1, 2, 3}, elements)
	assert.True(t, ok)
}

func TestIntDoubleLinkedList_PushFront(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	elements := []int{}
	value := 0
	ok := false

	// Empty Data
	value, ok = sut.Head()
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = sut.Tail()
	assert.Zero(t, value)
	assert.False(t, ok)

	assert.Equal(t, 0, sut.Length())
	elements, ok = sut.GetThToLast(0)
	assert.Nil(t, elements)
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
	elements, ok = sut.GetThToLast(0)
	assert.Equal(t, []int{1}, elements)
	assert.True(t, ok)

	// Push Front 2
	sut.PushFront(2)

	value, ok = sut.Head()
	assert.Equal(t, 2, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	assert.Equal(t, 2, sut.Length())
	elements, ok = sut.GetThToLast(0)
	assert.Equal(t, []int{2, 1}, elements)
	assert.True(t, ok)

	// Push Front 3
	sut.PushFront(3)

	value, ok = sut.Head()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	assert.Equal(t, 3, sut.Length())
	elements, ok = sut.GetThToLast(0)
	assert.Equal(t, []int{3, 2, 1}, elements)
	assert.True(t, ok)
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

// Cracking The Code Interview
func TestIntDoubleLinkedList_RemoveDuplication(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	sut.PushBack(1)
	sut.PushBack(1)
	sut.PushBack(2)
	sut.PushBack(1)
	sut.PushBack(2)
	sut.PushBack(3)
	sut.PushBack(2)
	sut.PushBack(2)

	// Act
	removed := sut.RemoveDuplication()
	value := 0
	ok := false

	// Assert
	value, ok = sut.Head()
	assert.Equal(t, 1, value)
	assert.True(t, ok)

	value, ok = sut.Tail()
	assert.Equal(t, 3, value)
	assert.True(t, ok)

	assert.Equal(t, 3, sut.Length())
	assert.Equal(t, []int{1, 1, 2, 2, 2}, removed)
}

func TestIntDoubleLinkedList_GetThToLast(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	for i := 0; i < 10; i++ {
		sut.PushBack(i)
	}

	// Act
	elements, ok := sut.GetThToLast(5)

	// Assert
	assert.True(t, ok)
	assert.Equal(t, []int{5, 6, 7, 8, 9}, elements)
}

func TestIntDoubleLinkedList_PartitionBetweenValue(t *testing.T) {
	// Arrange
	sut := data.NewDoubleIntLinkedList()

	for i := 0; i < 10; i++ {
		sut.PushBack(10 - i)
	}

	// Act
	sut, ok := sut.PartitionBetweenValue(5)

	// Assert
	assert.True(t, ok)
	assert.Equal(t, 10, sut.Length())
	elements, ok := sut.GetThToLast(0)
	assert.True(t, ok)
	assert.Equal(t, []int{4, 3, 2, 1, 10, 9, 8, 7, 6, 5}, elements)
}
