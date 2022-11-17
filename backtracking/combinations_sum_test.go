package backtracking

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputsCombinationsSumArrs = [][]int{{2, 3}, {2, 3, 6, 7}}
var inputsCombinationSums = []int{2, 7}
var outputsCombinationsSum = [][][]int{{{2}}, {{2, 2, 3}, {7}}}

func TestGetCombinationsWithSum(t *testing.T) {

	assert := assert.New(t)

	for i, inputCombinationsSum := range inputsCombinationSums {

		result := getCombinationsWithSum(inputsCombinationsSumArrs[i], inputCombinationsSum)
		assert.True(reflect.DeepEqual(result, outputsCombinationsSum[i]), fmt.Sprintf("Didn't get expected output for inputs: %v, %d\n Expected:%v, Got: %v ", inputsCombinationsSumArrs[i], inputCombinationsSum, outputsCombinationsSum[i], result))
	}
}
