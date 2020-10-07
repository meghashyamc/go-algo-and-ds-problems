package trees

import (
	"reflect"
	"testing"
)

var t0, t1 *treeNode

func setup() {
	t0rightleft := &treeNode{value: 15}
	t0rightright := &treeNode{value: 7}
	t0left := &treeNode{value: 9}
	t0right := &treeNode{left: t0rightleft, right: t0rightright, value: 20}
	t0 = &treeNode{left: t0left, right: t0right, value: 3}
	inputsReverseLevel = append(inputsReverseLevel, t0)

	t1rightleft := &treeNode{value: 3}
	t1left := &treeNode{value: 6}
	t1right := &treeNode{left: t1rightleft, value: 2}
	t1 = &treeNode{left: t1left, right: t1right, value: 1}
	inputsReverseLevel = append(inputsReverseLevel, t1)

}

var inputsReverseLevel = make([]*treeNode, 0)
var outputsReverseLevel = [][]int{{15, 7, 9, 20, 3}, {3, 6, 2, 1}}

func TestSolve(t *testing.T) {
	setup()
	for i, input := range inputsReverseLevel {
		res := solve(input)
		if !reflect.DeepEqual(outputsReverseLevel[i], res) {
			t.Error("expected output", outputsReverseLevel[i], "for", input, "got", res)
		}
	}
}
