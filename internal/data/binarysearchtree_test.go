package data_test

import (
	"algorithm-study/internal/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBinaryIntSearchTree(t *testing.T) {
	// Act
	sut := data.NewBinaryIntSearchTree()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, 0, sut.Length())
}

func TestBinaryIntSearchTree_Add(t *testing.T) {
	// Arrange
	sut := data.NewBinaryIntSearchTree()

	value := 0
	ok := true

	// Zero Adds
	value, ok = sut.Root()
	assert.Zero(t, value)
	assert.False(t, ok)
	assert.Equal(t, 0, sut.Length())

	// First Add
	sut.Add(100)
	value, ok = sut.Root()
	assert.Equal(t, 100, value)
	assert.True(t, ok)
	assert.Equal(t, 1, sut.Length())

	// Second Add
	sut.Add(200)
	value, ok = sut.Root()
	assert.Equal(t, 100, value)
	assert.True(t, ok)
	assert.Equal(t, 2, sut.Length())
}

func TestBinaryIntSearchTree_Exists(t *testing.T) {
	// Arrange
	sut := data.NewBinaryIntSearchTree()

	sut.Add(500)
	sut.Add(250)
	sut.Add(1000)

	exists := false

	// Doesnt Exists
	exists = sut.Exists(666)
	assert.False(t, exists)

	// Exists Root
	exists = sut.Exists(500)
	assert.True(t, exists)

	// Exists Left
	exists = sut.Exists(250)
	assert.True(t, exists)

	// Exists Right
	exists = sut.Exists(1000)
	assert.True(t, exists)
}
