package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var segregateZeroesAndOnesInputs = [][]int{{0, 1, 0}, {1, 1, 0}}
var segregateZeroesAndOnesOutputs = [][]int{{0, 0, 1}, {0, 1, 1}}

func TestSegregateZeroesAndOnes(t *testing.T) {

	assert := require.New(t)
	for i, input := range segregateZeroesAndOnesInputs {
		assert.Equal(segregateZeroesAndOnesOutputs[i], segregateZeroesAndOnes(input), "output array not as expected")
	}

}
