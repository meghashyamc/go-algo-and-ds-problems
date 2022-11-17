package trees

/*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).



Input Format
First and only argument is root node of the binary tree, A.



Output Format
Return a 2D integer array denoting the zigzag level order traversal of the given binary tree.



Example Input
Input 1:

    3
   / \
  9  20
    /  \
   15   7
Input 2:

   1
  / \
 6   2
    /
   3


Example Output
Output 1:

 [
   [3],
   [9, 20],
   [15, 7]
 ]
Output 2:

 [
   [1]
   [6, 2]
   [3]
 ]


Example Explanation
Explanation 1:

 Return the 2D array. Each row denotes the traversal of each level. */

func levelOrder(node *treeNode) [][]int {

	levelOrderArr := [][]int{}

	if node == nil {
		return levelOrderArr
	}

	q := queue{}

	q.rpush(&treeNodeWithLevel{node: node, level: 0})
	for {

		if len(q) == 0 {
			break
		}
		currNode := q.lpop()
		if len(levelOrderArr) >= currNode.level+1 {
			levelOrderArr[currNode.level] = append(levelOrderArr[currNode.level], currNode.node.value)
		} else {
			levelOrderArr = append(levelOrderArr, []int{currNode.node.value})
		}
		if currNode.node.left != nil {
			q.rpush(&treeNodeWithLevel{node: currNode.node.left, level: currNode.level + 1})
		}
		if currNode.node.right != nil {
			q.rpush(&treeNodeWithLevel{node: currNode.node.right, level: currNode.level + 1})
		}

	}

	return levelOrderArr

}
