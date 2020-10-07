package heapsandmaps

import (
	"reflect"
	"testing"
)

var inputsKLargest = [][]int{{11, 3, 4}, {11, 3, 4, 6}}
var inputsB = []int{2, 3}
var outputsKLargest = [][]int{{11, 4}, {11, 6, 4}}

func TestSolve(t *testing.T) {

	for i, input := range inputsKLargest {

		res := solveHeap(input, inputsB[i])
		if !reflect.DeepEqual(res, outputsKLargest[i]) {
			t.Error("expected", outputsKLargest[i], "got", res, "for input array", input)
		}
	}

}
