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

	sut.Add(2)
	sut.Add(3)
	sut.Add(4)
	sut.Add(5)
	sut.Add(6)
	sut.Add(7)
	sut.Add(8)
	sut.Add(9)
	sut.Print()
}
