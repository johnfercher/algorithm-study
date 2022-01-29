package data_test

import (
	"algorithm-study/internal/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueueIntLinkedList(t *testing.T) {
	// Act
	sut := data.NewQueueIntLinkedList()

	// Assert
	assert.Zero(t, sut.Length())
	value, ok := sut.First()
	assert.Zero(t, value)
	assert.False(t, ok)
}

func TestQueueIntLinkedList_Enqueue(t *testing.T) {
	// Arrange
	sut := data.NewQueueIntLinkedList()

	value := 0
	ok := false

	// Enqueue 1
	sut.Enqueue(1)
	value, ok = sut.First()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	value, ok = sut.Last()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	assert.Equal(t, 1, sut.Length())

	// Enqueue 2
	sut.Enqueue(2)
	value, ok = sut.First()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	value, ok = sut.Last()
	assert.Equal(t, 2, value)
	assert.True(t, ok)
	assert.Equal(t, 2, sut.Length())

	// Enqueue 3
	sut.Enqueue(3)
	value, ok = sut.First()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	value, ok = sut.Last()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
	assert.Equal(t, 3, sut.Length())
}

func TestQueueIntLinkedList_Dequeue(t *testing.T) {
	// Arrange
	sut := data.NewQueueIntLinkedList()

	value := 0
	ok := false

	sut.Enqueue(1)
	sut.Enqueue(2)
	sut.Enqueue(3)

	sut.Print()

	// Dequeue 1
	value, ok = sut.Dequeue()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	value, ok = sut.First()
	assert.Equal(t, 2, value)
	assert.True(t, ok)
	value, ok = sut.Last()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
	assert.Equal(t, 2, sut.Length())

	// Dequeue 2
	value, ok = sut.Dequeue()
	assert.Equal(t, 2, value)
	assert.True(t, ok)
	value, ok = sut.First()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
	value, ok = sut.Last()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
	assert.Equal(t, 1, sut.Length())

	// Dequeue 3
	value, ok = sut.Dequeue()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
	value, ok = sut.First()
	assert.Zero(t, value)
	assert.False(t, ok)
	value, ok = sut.Last()
	assert.Zero(t, value)
	assert.False(t, ok)
	assert.Equal(t, 0, sut.Length())

	// Invalid
	value, ok = sut.Dequeue()
	assert.Zero(t, value)
	assert.False(t, ok)
	value, ok = sut.First()
	assert.Zero(t, value)
	assert.False(t, ok)
	value, ok = sut.Last()
	assert.Zero(t, value)
	assert.False(t, ok)
	assert.Equal(t, 0, sut.Length())
}

func TestQueueIntLinkedList_Peek(t *testing.T) {
	// Arrange
	sut := data.NewQueueIntLinkedList()

	value := 0
	ok := false

	sut.Enqueue(1)
	sut.Enqueue(2)
	sut.Enqueue(3)

	// Enqueue 1
	value, ok = sut.First()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
}
