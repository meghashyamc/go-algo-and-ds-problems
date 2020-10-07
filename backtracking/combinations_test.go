package backtracking

import (
	"reflect"
	"testing"
)

var inputsCombineN = []int{4}
var inputsCombineK = []int{2}
var outputsCombine = [][][]int{{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}}

func TestCombine(t *testing.T) {

	for i := 0; i < len(inputsCombineK); i++ {
		res := combine(inputsCombineN[i], inputsCombineK[i])
		if !reflect.DeepEqual(res, outputsCombine[i]) {
			t.Error("expected combinations to be", outputsCombine[i], "for", inputsCombineN[i], "and", inputsCombineK[i], "got", res)
		}
	}
}
