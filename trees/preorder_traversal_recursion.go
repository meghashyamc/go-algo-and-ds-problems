package trees

/* Given a binary tree, return the preorder traversal of its nodes’ values.

Example :

Given binary tree

   1
    \
     2
    /
   3
return [1,2,3]. */

func preOrderTraversalRecursion(node *treeNode) []int {

	preOrderArr := []int{}
	if node == nil {
		return preOrderArr
	}

	preOrderTraversalRecursionHelper(node, &preOrderArr)

	return preOrderArr
}

func preOrderTraversalRecursionHelper(node *treeNode, preOrderArr *[]int) {

	if node == nil {
		return
	}

	*preOrderArr = append(*preOrderArr, node.value)
	preOrderTraversalRecursionHelper(node.left, preOrderArr)
	preOrderTraversalRecursionHelper(node.right, preOrderArr)

}
