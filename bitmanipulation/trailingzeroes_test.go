package bitmanipulation

import (
	"testing"
)

var inputTrailingZeroes = []int{18, 8, 4}
var outputTrailingZeroes = []int{1, 3, 2}

func TestSolve(t *testing.T) {

	for i, input := range inputTrailingZeroes {
		res := solve(input)
		if res != outputTrailingZeroes[i] {
			t.Error("expected number of zeroes to be", outputTrailingZeroes[i], "got input", input, "got", res)
		}
	}
}
