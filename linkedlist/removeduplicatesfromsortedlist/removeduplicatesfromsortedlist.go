package removeduplicatesfromsortedlist

import ln "github.com/wenndersantos/leetcode-in-go/linkedlist/listnode"

//DeleteDuplicates executes in
//Time: O(n) where n is the length of the linkedlist
//Memory: O(1)
func DeleteDuplicates(head *ln.ListNode) *ln.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

//83. Remove Duplicates from Sorted List
//https://leetcode.com/problems/remove-duplicates-from-sorted-list/
