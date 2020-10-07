package stacksandqueues

import (
	"testing"
)

var inputsRedundantBraces = []string{"((a+b))", "(a+(a+b))", "(a*b)+(b*(d+(a)))", "(a*b)"}
var outputsRedundantBraces = []int{1, 0, 1, 0}

func TestHasRedundantBraces(t *testing.T) {

	for i, input := range inputsRedundantBraces {
		res := hasRedundantBraces(input)
		if res != outputsRedundantBraces[i] {
			t.Error("expected", outputsRedundantBraces[i], "for input", input, "got", res)
		}
	}
}
