package arrays

import (
	"reflect"
	"testing"
)

var inputArrays = [][]int{{-6, -3, -1, 2, 4, 5}, {-5, -4, -2, 0, 1}}
var outputArrays = [][]int{{1, 4, 9, 16, 25, 36}, {0, 1, 4, 16, 25}}

func TestSortArrayWithSquares(t *testing.T) {

	for i, input := range inputArrays {

		res := sortArrayWithSquares(input)
		if !reflect.DeepEqual(outputArrays[i], res) {
			t.Error("expected output", outputArrays[i], "for input", input, "got", res)
		}
	}

}
