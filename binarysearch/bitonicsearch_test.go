package binarysearch

import "testing"

var inputsBitonicSearch = [][]int{{3, 9, 10, 20, 17, 5, 1, 20}, {5, 6, 7, 8, 9, 10, 3, 2, 1, 30}}
var outputsBitonicSearch = []int{3, -1}

func TestSolve(t *testing.T) {

	for i, input := range inputsBitonicSearch {

		res := solve(input[:len(input)-1], input[len(input)-1])
		if res != outputsBitonicSearch[i] {
			t.Fatal("expected output", outputsBitonicSearch[i], "for input", input, "got", res)
		}
	}
}
