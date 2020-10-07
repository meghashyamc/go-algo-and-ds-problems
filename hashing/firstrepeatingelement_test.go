package hashing

import (
	"testing"
)

var inputsRepeatingElement = [][]int{{10, 5, 3, 4, 3, 5, 6}}
var outputsRepeatingElement = []int{5}

func TestSolve(t *testing.T) {

	for i, input := range inputsRepeatingElement {
		res := solve(input)
		if res != outputsRepeatingElement[i] {
			t.Error("expected output", outputsRepeatingElement[i], "for input", input, "got", res)
		}
	}
}
