package math

import (
	"testing"
)

var inputsTrailingZeroes = []int{4, 5, 25, 125}
var outputsTrailingZeroes = []int{0, 1, 6, 31}

func TestTrailingZeroes(t *testing.T) {

	for i, input := range inputsTrailingZeroes {
		res := trailingZeroes(input)
		if res != outputsTrailingZeroes[i] {
			t.Error("expected trailing zeroes to be", outputsTrailingZeroes[i], "for", input, "got", res)
		}

	}

}
