package data_test

import (
	"algorithm-study/internal/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIntHashSet(t *testing.T) {
	// Act
	sut := data.NewIntHashSet(1000)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, 0, sut.Length())
}

func TestIntHashSet_Add(t *testing.T) {
	// Arrange
	sut := data.NewIntHashSet(1000)

	// Act
	sut.Add("key", 10)

	// Assert
	value, ok := sut.Get("key")
	assert.Equal(t, 1, sut.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, 10, value)
}

func TestIntHashSet_Remove(t *testing.T) {
	// Arrange
	sut := data.NewIntHashSet(1000)

	// Act
	sut.Add("key", 10)
	sut.Remove("key")

	// Assert
	value, ok := sut.Get("key")
	assert.Equal(t, 0, sut.Length())
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, value)
}
