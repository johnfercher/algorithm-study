package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStackIntLinkedListMin(t *testing.T) {
	// Act
	sut := NewStackIntLinkedListMin()

	// Assert
	assert.NotNil(t, sut)
}

func TestStackIntLinkedListMin_Min(t *testing.T) {
	// Arrange
	sut := NewStackIntLinkedListMin()

	value := 0
	ok := false

	// Act & Assert
	value, ok = sut.Min()
	assert.Equal(t, 0, value)
	assert.False(t, ok)
	assert.Equal(t, 0, sut.Length())

	sut.Push(4)
	value, ok = sut.Min()
	assert.Equal(t, 4, value)
	assert.True(t, ok)
	assert.Equal(t, 1, sut.Length())

	sut.Push(3)
	value, ok = sut.Min()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
	assert.Equal(t, 2, sut.Length())

	sut.Push(2)
	value, ok = sut.Min()
	assert.Equal(t, 2, value)
	assert.True(t, ok)
	assert.Equal(t, 3, sut.Length())

	sut.Push(1)
	value, ok = sut.Min()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	assert.Equal(t, 4, sut.Length())

	sut.Pop()
	value, ok = sut.Min()
	assert.Equal(t, 2, value)
	assert.True(t, ok)
	assert.Equal(t, 3, sut.Length())

	sut.Pop()
	value, ok = sut.Min()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
	assert.Equal(t, 2, sut.Length())

	sut.Pop()
	value, ok = sut.Min()
	assert.Equal(t, 4, value)
	assert.True(t, ok)
	assert.Equal(t, 1, sut.Length())
}
