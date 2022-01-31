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

	// Act
	sut.Add(1)

	sut.Add(3)
	sut.Add(2)
	sut.Print()
}
