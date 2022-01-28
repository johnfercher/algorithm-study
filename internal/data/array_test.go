package data_test

import (
	"algorithm-study/internal/data"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	// Act
	sut := data.NewIntArray()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*data.IntArray", fmt.Sprintf("%T", sut))
	assert.Zero(t, sut.Len())
}

func TestIntArray_Push(t *testing.T) {
	// Arrange
	sut := data.NewIntArray()

	// Act
	sut.Push(0)
	sut.Push(10)

	// Assert
	assert.Equal(t, 2, sut.Len())
	assert.Equal(t, 10, sut.At(sut.Len()-1))
}

func TestIntArray_PushFront(t *testing.T) {
	// Arrange
	sut := data.NewIntArray()

	// Act
	sut.Push(0)
	sut.PushFront(10)

	// Assert
	assert.Equal(t, 2, sut.Len())
	assert.Equal(t, 10, sut.At(0))
}

func TestIntArray_Pop(t *testing.T) {
	// Arrange
	sut := data.NewIntArray()

	// Act
	sut.Push(0)
	sut.Push(10)
	sut.Print()
	removedValue := sut.Pop()

	// Assert
	assert.Equal(t, 1, sut.Len())
	assert.Equal(t, 10, removedValue)
	sut.Print()
}

func TestIntArray_PopFront(t *testing.T) {
	// Arrange
	sut := data.NewIntArray()

	// Act
	sut.Push(10)
	sut.Push(0)
	removedValue := sut.PopFront()

	// Assert
	assert.Equal(t, 1, sut.Len())
	assert.Equal(t, 10, removedValue)
}
