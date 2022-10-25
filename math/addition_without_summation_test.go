package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputsAddNumbersWithoutOperator = [][]int{{761, 384}}
var outputsAddNumbersWithoutOperator = []int{1145}

func TestAddNumbersWithoutOperator(t *testing.T) {

	assert := assert.New(t)

	for i, input := range inputsAddNumbersWithoutOperator {
		gottenOutput := addNumbersWithoutOperator(input[0], input[1])
		assert.Equal(outputsAddNumbersWithoutOperator[i], gottenOutput, fmt.Sprintf("expected output not matching gotten output for inputs %d and %d", input[0], input[1]))
	}
}
