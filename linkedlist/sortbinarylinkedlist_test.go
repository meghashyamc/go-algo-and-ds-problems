package linkedlist

import (
	"reflect"
	"testing"
)

var inputLists = [][]int{{1, 0, 0, 1}, {0, 0, 1, 1, 0}, {1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0}}
var outputLists = [][]int{{0, 0, 1, 1}, {0, 0, 0, 1, 1}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}

func setup(arr []int) *listNode {
	var head *listNode
	prevnode := &listNode{value: arr[len(arr)-1]}
	for i := len(arr) - 2; i >= 0; i-- {
		currnode := &listNode{value: arr[i], next: prevnode}
		if i != 0 {
			prevnode = currnode
		} else {
			head = currnode

		}
	}

	return head

}

func TestSolve(t *testing.T) {

	for i, input := range inputLists {

		res := solve(setup(input))
		arrayRes := convertToArray(res)
		if !reflect.DeepEqual(arrayRes, outputLists[i]) {
			t.Error("expected output", outputLists[i], "for input", input, "got", arrayRes)
		}
	}

}

func convertToArray(node *listNode) []int {

	curr := node

	arr := make([]int, 0)

	for curr != nil {

		arr = append(arr, curr.value)
		curr = curr.next
	}
	return arr
}
