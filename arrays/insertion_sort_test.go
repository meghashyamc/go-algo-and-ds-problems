package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputsInsertionSort = [][]int{{-3, -6, -1, 4, 5, 2}, {0, 1, -2, -5, -4}}
var outputsInsertionSort = [][]int{{-6, -3, -1, 2, 4, 5}, {-5, -4, -2, 0, 1}}

func TestInsertionSort(t *testing.T) {

	assert := assert.New(t)

	for i, input := range inputsInsertionSort {
		inputCopy := make([]int, len(input))
		copy(inputCopy, input)
		insertionSort(inputCopy)
		assert.Equal(outputsInsertionSort[i], inputCopy)
	}
}
