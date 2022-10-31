package linkedlist

/* Given a linked list, remove the nth node from the end of list and return its head.

For example,

   Given linked list: 1->2->3->4->5, and n = 2.

   After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:

If n is greater than the size of the list, remove the first node of the list.
Try doing it using constant additional space. */

func removeNthNodeFromEnd(head *listNode, n int) *listNode {

	if n <= 0 {
		return head
	}
	nodesCount := getNodesCountForLinkedList(head)

	indexOfNodeBeforeNodeToRemove := nodesCount - n - 1

	if indexOfNodeBeforeNodeToRemove < 0 {
		newHead := head.next
		head.next = nil
		return newHead
	}
	currNode := head
	var nodeBeforeNodeToRemove *listNode
	for i := 0; i < nodesCount; i++ {
		if i == indexOfNodeBeforeNodeToRemove {
			nodeBeforeNodeToRemove = currNode

		} else if i == indexOfNodeBeforeNodeToRemove+1 {
			nodeBeforeNodeToRemove.next = currNode.next
			currNode.next = nil
			break

		}
		currNode = currNode.next

	}

	return head

}
