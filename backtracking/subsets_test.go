package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

var inputsSubsets = [][]int{{1, 2, 3}, {15, 20, 12, 19, 4}}
var outputsSubsets = [][][]int{{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}, {{}, {4}, {4, 12}, {4, 12, 15}, {4, 12, 15, 19}, {4, 12, 15, 19, 20}, {4, 12, 15, 20}, {4, 12, 19}, {4, 12, 19, 20}, {4, 12, 20}, {4, 15}, {4, 15, 19}, {4, 15, 19, 20}, {4, 15, 20}, {4, 19}, {4, 19, 20}, {4, 20}, {12}, {12, 15}, {12, 15, 19}, {12, 15, 19, 20}, {12, 15, 20}, {12, 19}, {12, 19, 20}, {12, 20}, {15}, {15, 19}, {15, 19, 20}, {15, 20}, {19}, {19, 20}, {20}}}

func TestSubsets(t *testing.T) {

	for i, input := range inputsSubsets {
		res := subsets(input)
		if !reflect.DeepEqual(res, outputsSubsets[i]) {
			t.Error("expected subsets", outputsSubsets[i], "for the set", input, "got", res)
		}
	}

}

func TestResultSort(t *testing.T) {
	arr := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
	sort.Slice(arr, func(i, j int) bool {
		return compareElements(arr, i, j)
	})
	sortedArr := [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}
	if len(arr) != len(sortedArr) {
		t.Fatal("expected length of sorted array to be", len(sortedArr), "got", len(arr))
	}
	if !reflect.DeepEqual(arr, sortedArr) {
		t.Fatal("expected", sortedArr, "got", arr)
	}
}
