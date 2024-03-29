package onstrings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	inputsVersionComparisonVersion1 = []string{"70809", "1.0"}
	inputsVersionComparisonVersion2 = []string{"9051", "1"}
	outputsVersionComparison        = []int{1, 0}
)

func TestCompareVersionNumbers(t *testing.T) {
	assert := require.New(t)
	for i := range inputsVersionComparisonVersion1 {

		assert.Equal(outputsVersionComparison[i], compareVersionNumbers(inputsVersionComparisonVersion1[i], inputsVersionComparisonVersion2[i]))

	}
}
