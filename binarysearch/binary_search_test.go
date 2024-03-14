package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputArrsBinarySearch = [][]int{{1, 2, 3, 4, 5}, {2, 10, 100, 100, 1001}, {-5, -2, 6, 9, 100}, {1, 50, 50, 50, 100}}
var inputNumsBinarySearch = []int{3, 100, 101, 50}

var outputsBinarySearch = [][]int{{2}, {2, 3}, {}, {1, 2, 3}}

func TestSearch(t *testing.T) {

	assert := assert.New(t)

	for i, inputArr := range inputArrsBinarySearch {
		assert.ElementsMatch(search(inputArr, inputNumsBinarySearch[i]), outputsBinarySearch[i], "error for input array %v and number being searched for %d", inputArr, inputNumsBinarySearch[i])
	}

}
