package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputsMergeSort = [][]int{{3, 41, 52, 26, 38, 57, 9, 49}, {-3, -6, -1, 4, 5, 2}, {0, 1, -2, -5, -4}}
var outputsMergeSort = [][]int{{3, 9, 26, 38, 41, 49, 52, 57}, {-6, -3, -1, 2, 4, 5}, {-5, -4, -2, 0, 1}}

func TestMergeSort(t *testing.T) {

	assert := assert.New(t)

	for i, input := range inputsMergeSort {
		inputCopy := make([]int, len(input))
		copy(inputCopy, input)
		mergeSort(inputCopy)
		assert.Equal(outputsMergeSort[i], inputCopy)
	}
}
