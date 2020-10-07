package stacksandqueues

import (
	"testing"
)

var inputParantheses = []string{"(()())", "(()", "", "()))))))))("}
var outputParantheses = []int{1, 0, 1, 0}

func TestSolve(t *testing.T) {

	for i, input := range inputParantheses {
		res := solve(input)
		if res != outputParantheses[i] {
			t.Error("expected", outputParantheses[i], "for input string", input, "got", res)
		}
	}
}
