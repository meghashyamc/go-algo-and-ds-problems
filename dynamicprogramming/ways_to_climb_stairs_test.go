package dynamicprogramming

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	inputsNumberOfWaysToClimbStairs = []int{2, 3}

	outputsNumberOfWaysToClimbStairs = []int{2, 3}
)

func TestGetNumberOfWaysToClimbStairs(t *testing.T) {

	assert := require.New(t)

	for i := 0; i < len(inputsNumberOfWaysToClimbStairs); i++ {

		assert.Equal(outputsNumberOfWaysToClimbStairs[i], getNumberOfWaysToClimbStairs(inputsNumberOfWaysToClimbStairs[i]), fmt.Sprintf("unexpected output when input is %d", inputsNumberOfWaysToClimbStairs[i]))
	}
}
