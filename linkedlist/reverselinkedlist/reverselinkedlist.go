package reverselinkedlist

import ln "github.com/wenndersantos/leetcode-in-go/linkedlist/listnode"

//ReverseList executes in
//Time: O(n) where n is the length of de list
//Memory: O(1) reversed in place
func ReverseList(head *ln.ListNode) *ln.ListNode {
	if head == nil {
		return head
	}

	current := head.Next
	head.Next = nil

	for current != nil {
		aux := current.Next
		current.Next = head

		head = current
		current = aux
	}

	return head
}

//206. Reverse Linked List
//https://leetcode.com/problems/reverse-linked-list/
