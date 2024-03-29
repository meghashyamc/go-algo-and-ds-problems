package hashing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var inputsLongestConsecutiveSequence = [][]int{{100, 4, 200, 1, 3, 2}}
var outputsLongestConsecutiveSequence = []int{4}

func TestGetLengthOfLongestConsecutiveSequence(t *testing.T) {

	assert := require.New(t)

	for i := 0; i < len(inputsLongestConsecutiveSequence); i++ {

		assert.Equal(outputsLongestConsecutiveSequence[i], getLengthOfLongestConsecutiveSequence(inputsLongestConsecutiveSequence[i]), fmt.Sprintf("got unexpected output for input %v", inputsLongestConsecutiveSequence[i]))
	}

}
