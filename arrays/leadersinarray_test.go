package arrays

import (
	"reflect"
	"testing"
)

var inputsSolve = [][]int{{16, 17, 4, 3, 5, 2}, {1, 2}}
var outputsSolve = [][]int{{2, 5, 17}, {2}}

func TestSolve(t *testing.T) {

	for i, input := range inputsSolve {
		res := solve(input)
		if !reflect.DeepEqual(res, outputsSolve[i]) {
			t.Error("expected output", outputsSolve[i], "for input", input, "got", res)
		}
	}
}
