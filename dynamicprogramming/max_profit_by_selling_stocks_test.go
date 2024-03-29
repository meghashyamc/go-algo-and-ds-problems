package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	inputsStockPrices = [][]int{{1, 2}}
	outputsMaxProfits = []int{1}
)

func TestGetMaxProfitBySellingStocks(t *testing.T) {

	assert := require.New(t)
	for i := 0; i < len(inputsStockPrices); i++ {

		assert.Equal(outputsMaxProfits[i], getMaxProfitBySellingStocks(inputsStockPrices[i]))
	}
}
