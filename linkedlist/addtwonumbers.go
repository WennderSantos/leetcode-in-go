package addtwonumbers

//ListNode is a node of a linkedlist
type ListNode struct {
	Val  int
	Next *ListNode
}

//AddTwoNumbers Executes in
//Time: O(n) where n is the length of the biggest list
//Memory: O(n) where n is the length of result list
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{}
	current := &head
	carry, sum := 0, 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum = carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return head.Next
}

//2. Add Two Numbers
//https://leetcode.com/problems/add-two-numbers/
