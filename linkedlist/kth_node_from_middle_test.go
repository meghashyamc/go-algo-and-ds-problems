package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputListsKthNodeFromMiddle = [][]int{{725, 479, 359, 963, 465, 706, 146, 282, 828, 962}}
var inputKsKthNodeFromMiddle = []int{2}
var outputsKthNodeFromMiddle = []int{963}

func getLinkedListFromArr(arr []int) *listNode {
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

func TestGetKthNodeFromMiddle(t *testing.T) {

	assert := assert.New(t)
	for i := 0; i < len(inputKsKthNodeFromMiddle); i++ {
		assert.Equal(outputsKthNodeFromMiddle[i], getKthNodeFromMiddle(getLinkedListFromArr(inputListsKthNodeFromMiddle[i]), inputKsKthNodeFromMiddle[i]))
	}

}
