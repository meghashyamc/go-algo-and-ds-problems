package trees

import "errors"

/* Given a binary tree, return the preorder traversal of its nodes’ values.

Example :

Given binary tree

   1
    \
     2
    /
   3
return [1,2,3].

DO NOT USE RECURSION
*/

type simpleStack []treeNode

func (s *simpleStack) push(val *treeNode) {

	if val == nil {
		return
	}
	*s = append(*s, *val)
}

func (s *simpleStack) pop() (*treeNode, error) {

	if len(*s) == 0 {
		return nil, errors.New("cannot pop empty stack")
	}
	ret := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]
	return &ret, nil
}
func preOrderTraversal(node *treeNode) []int {

	preOrderArr := []int{}
	if node == nil {
		return preOrderArr
	}

	stack := simpleStack{}

	currNode := node

	for {

		if currNode != nil {
			stack.push(currNode)
			preOrderArr = append(preOrderArr, node.value)
			currNode = currNode.left
			continue
		}

		parentNode, err := stack.pop()
		if err != nil {
			break
		}
		currNode = parentNode.right

	}

	return preOrderArr
}
