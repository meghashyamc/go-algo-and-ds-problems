package strings

import (
	"testing"
)

var inputStrings = []string{"aabcd", "aabbccd"}
var inputBs = []int{2, 2}
var outputStrings = []string{"bcd", "d"}

func TestSolve(t *testing.T) {

	for i := 0; i < len(inputStrings); i++ {
		res := solve(inputStrings[i], inputBs[i])
		if res != outputStrings[i] {
			t.Error("expected", outputStrings[i], "for input string", inputStrings[i], "got", res)
		}
	}
}
