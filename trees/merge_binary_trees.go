package trees

/* Given two Binary Trees A and B, you need to merge them in a single binary tree.

The merge rule is that if two nodes overlap, then sum of node values is the new value of the merged node.

Otherwise, the non-null node will be used as the node of new tree.



Problem Constraints
1 <= Number of Nodes in A , B <= 10^5



Input Format
First argument is a pointer to the root of binary tree A.

Second argument is a pointer to the root of binary tree B.



Output Format
Return a pointer to the root of new binary tree.



Example Input
Input 1:

 A =   2

      / \

     1   4

    /

   5


B =   3
      / \
      6  1
      \   \
       2   7

Input 2:

 A =   12

      / \

     11  14


B =   3
      / \
      6  1



Example Output
Output 1:

     5
    / \
   7   5
  / \   \
 5   2   7
Output 2:

   15
  / \
 17  15


Example Explanation
Explanation 1:

 After merging both the trees you get:
     5
    / \
   7   5
  / \   \
 5   2   7
Explanation 2:

 After merging both the trees we get:
   15
  / \
 17  15 */

func mergeBinaryTrees(node1 *treeNode, node2 *treeNode) *treeNode {

	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	mergedNode := &treeNode{value: node1.value + node2.value}

	mergedNode.left = mergeBinaryTrees(node1.left, node2.left)
	mergedNode.right = mergeBinaryTrees(node1.right, node2.right)

	return mergedNode

}
