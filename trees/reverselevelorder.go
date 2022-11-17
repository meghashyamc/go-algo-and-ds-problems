package trees

import (
	"log"
)

// 	Problem Description

// Given a binary tree, return the reverse level order traversal of its nodes values. (i.e, from left to right and from last level to starting level).

// Problem Constraints
// 1 <= number of nodes <= 5 * 105

// 1 <= node value <= 105

// Input Format
// First and only argument is a pointer to the root node of the Binary Tree, A.

// Output Format
// Return an integer array denoting the reverse level order traversal from left to right.

// Example Input
// Input 1:

//     3
//    / \
//   9  20
//     /  \
//    15   7
// Input 2:

//    1
//   / \
//  6   2
//     /
//    3

// Example Output
// Output 1:

//  [15, 7, 9, 20, 3]
// Output 2:

//  [3, 6, 2, 1]

// Example Explanation
// Explanation 1:

//  Nodes as level 3 : [15, 7]
//  Nodes at level 2: [9, 20]
//  Nodes at level 1: [3]
//  Reverse level order traversal will be: [15, 7, 9, 20, 3]
// Explanation 2:

// Nodes as level 3 : [3]
// Nodes at level 2: [6, 2]
// Nodes at level 1: [1]
// Reverse level order traversal will be: [3, 6, 2, 1]
type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func treeNodeNew(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

type treeNodeWithLevel struct {
	node  *treeNode
	level int
}

type queue []*treeNodeWithLevel

func (q *queue) rpush(i *treeNodeWithLevel) {

	if q == nil {
		q = new(queue)
	}
	*q = append(*q, i)
}

func (q *queue) lpop() *treeNodeWithLevel {

	if q == nil || len(*q) == 0 {
		log.Fatal("empty queue cannot be popped")
	}
	popped := (*q)[0]
	*q = (*q)[1:]
	return popped

}
func solve(A *treeNode) []int {

	if A == nil {
		return []int{}
	}
	levelOrderArr := make([][]int, 0)
	q := queue{}
	q.rpush(&treeNodeWithLevel{node: A, level: 0})
	for len(q) > 0 {
		popped := q.lpop()

		if len(levelOrderArr) >= popped.level+1 {
			levelOrderArr[popped.level] = append(levelOrderArr[popped.level], popped.node.value)

		} else {
			newLevel := make([]int, 0)
			newLevel = append(newLevel, popped.node.value)
			levelOrderArr = append(levelOrderArr, newLevel)
		}
		if popped.node.left != nil {
			q.rpush(&treeNodeWithLevel{node: popped.node.left, level: popped.level + 1})
		}
		if popped.node.right != nil {
			q.rpush(&treeNodeWithLevel{node: popped.node.right, level: popped.level + 1})
		}
	}
	return getReverseLevelOrder(levelOrderArr)

}
func getReverseLevelOrder(arr [][]int) []int {

	reverseLevelArr := make([]int, 0)
	if arr == nil || len(arr) == 0 {
		return reverseLevelArr
	}

	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < len(arr[i]); j++ {
			reverseLevelArr = append(reverseLevelArr, arr[i][j])
		}
	}
	return reverseLevelArr

}
