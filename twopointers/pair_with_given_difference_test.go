package twopointers

import (
	"testing"
)

var inputsPairDifferenceArr = [][]int{{5, 10, 3, 2, 50, 80}, {-10, 20}, {478, 358, -38, -536, 705}, {-533, -666, -500, 169, 724, 478, 358, -38, -536, 705, -855, 281, -173, 961, -509, -5, 942, -173, 436, -609, -396, 902, -847, -708, -618, 421, -284, 718, 895, 447, 726, -229, 538, 869, 912, 667, -701, 35, 894, -297, 811, 322}}
var inputsB = []int{78, 30, 320, 369}
var outputsPairDifference = []int{1, 1, 0, 1}

func TestSolve(t *testing.T) {

	for i := 0; i < len(inputsB); i++ {
		res := solve(inputsPairDifferenceArr[i], inputsB[i])

		if res != outputsPairDifference[i] {
			t.Error("expected output", outputsPairDifference[i], "for input array", inputsPairDifferenceArr[i], "got", res)
		}
	}
}
