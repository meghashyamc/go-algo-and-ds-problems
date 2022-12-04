package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputsHeapSort = [][]int{{-3, -6, -1, 4, 5, 2}, {0, 1, -2, -5, -4}}
var outputsHeapSort = [][]int{{-6, -3, -1, 2, 4, 5}, {-5, -4, -2, 0, 1}}

func TestHeapSort(t *testing.T) {

	assert := assert.New(t)

	for i, input := range inputsHeapSort {
		inputCopy := make([]int, len(input))
		copy(inputCopy, input)
		heapSort(inputCopy)
		assert.Equal(outputsHeapSort[i], inputCopy)
	}
}
