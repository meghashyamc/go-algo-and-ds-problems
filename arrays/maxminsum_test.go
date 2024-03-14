package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInputsForMaxMinElementSum = [][]int{{-2, 1, -4, 5, 3}, {1, 3, 4, 1}}
var testOutputsForMaxMinElementSum = []int{1, 5}

func TestGetMaxMinElementSum(t *testing.T) {

	assert := require.New(t)

	for i, input := range testInputsForMaxMinElementSum {
		assert.Equal(testOutputsForMaxMinElementSum[i], getMaxMinElementSum(input))

	}
}
