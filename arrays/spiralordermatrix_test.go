package arrays

import (
	"reflect"
	"testing"
)

var inputsSpiralOrder = [][][]int{{{1}}, {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}
var outputsSpiralOrder = [][]int{{1}, {1, 2, 3, 6, 9, 8, 7, 4, 5}}

func TestSpiralOrder(t *testing.T) {

	for i, input := range inputsSpiralOrder {
		res := spiralOrder(input)
		if !reflect.DeepEqual(outputsSpiralOrder[i], res) {
			t.Fatal("expected output", outputsSpiralOrder[i], "for input", input, "got", res)
		}
	}
}
