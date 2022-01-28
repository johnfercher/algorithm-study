package hash_test

import (
	hash "algorithm-study/internal/hash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	// Arrange
	maxValue := 1000
	value := "hello"

	// Act
	hashValue := hash.Generate(value, maxValue)

	// Assert
	assert.Equal(t, 85, hashValue)
}
