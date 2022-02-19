package sort_test

import (
	"algorithm-study/internal/generate"
	"algorithm-study/internal/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionInt_WhenOrderAsc_ShouldOrderAsc(t *testing.T) {
	// Arrange
	unsortedArr := generate.UnsortedIntArray(10, 100)

	valuesMap := make(map[int]bool)
	for _, value := range unsortedArr {
		valuesMap[value] = true
	}

	// Act
	sortedArr := sort.SelectionInt(unsortedArr, false)

	// Assert
	assert.Equal(t, len(unsortedArr), len(sortedArr))
	for i := 0; i < len(sortedArr)-1; i++ {
		assert.LessOrEqual(t, sortedArr[i+1], sortedArr[i])
		assert.True(t, valuesMap[sortedArr[i]])
	}
}

func TestSelectionInt_WhenOrderDesc_ShouldOrderDesc(t *testing.T) {
	// Arrange
	unsortedArr := generate.UnsortedIntArray(10, 100)

	valuesMap := make(map[int]bool)
	for _, value := range unsortedArr {
		valuesMap[value] = true
	}

	// Act
	sortedArr := sort.SelectionInt(unsortedArr, false)

	// Assert
	assert.Equal(t, len(unsortedArr), len(sortedArr))
	for i := 0; i < len(sortedArr)-1; i++ {
		assert.LessOrEqual(t, sortedArr[i+1], sortedArr[i])
		assert.True(t, valuesMap[sortedArr[i]])
	}
}
