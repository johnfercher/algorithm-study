package datastructures_test

import (
	"algorithm-study/internal/datastructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueueWithStack(t *testing.T) {
	// Act
	sut := datastructures.NewQueueWithStack()

	// Assert
	assert.Zero(t, sut.Length())
	value, ok := sut.First()
	assert.Zero(t, value)
	assert.False(t, ok)
}

func TestQueueWithStack_Enqueue(t *testing.T) {
	// Arrange
	sut := datastructures.NewQueueWithStack()

	value := 0
	ok := false

	// Enqueue 1
	sut.Enqueue(1)
	value, ok = sut.First()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	assert.Equal(t, 1, sut.Length())

	// Enqueue 2
	sut.Enqueue(2)
	value, ok = sut.First()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	assert.Equal(t, 2, sut.Length())

	// Enqueue 3
	sut.Enqueue(3)
	value, ok = sut.First()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	assert.Equal(t, 3, sut.Length())
}

func TestQueueWithStack_Dequeue(t *testing.T) {
	// Arrange
	sut := datastructures.NewQueueWithStack()

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
	assert.Equal(t, 2, sut.Length())

	// Dequeue 2
	value, ok = sut.Dequeue()
	assert.Equal(t, 2, value)
	assert.True(t, ok)
	value, ok = sut.First()
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
	assert.Equal(t, 0, sut.Length())

	// Invalid
	value, ok = sut.Dequeue()
	assert.Zero(t, value)
	assert.False(t, ok)
	value, ok = sut.First()
	assert.Zero(t, value)
	assert.False(t, ok)
	assert.Equal(t, 0, sut.Length())
}

func TestQueueWithStack_Peek(t *testing.T) {
	// Arrange
	sut := datastructures.NewQueueWithStack()

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
