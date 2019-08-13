package designlinkedlist

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{Head: nil,
		Tail:   nil,
		Length: 0}
}

//Get the value of the index-th node in the linked list. If the index is invalid, return -1.
func (MyLinkedList *MyLinkedList) Get(index int) int {
	if index < 1 || index > MyLinkedList.Length {
		return -1
	}

	current := MyLinkedList.Head
	index--

	for index > 0 {
		current = current.Next
		index--
	}

	return current.Val
}

//AddAtHead a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (MyLinkedList *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{Val: val, Next: MyLinkedList.Head}
	MyLinkedList.Head = newNode
	MyLinkedList.Length++

	if MyLinkedList.Tail == nil {
		MyLinkedList.Tail = newNode
	}
}

//AddAtTail a node of value val to the last element of the linked list.
func (MyLinkedList *MyLinkedList) AddAtTail(val int) {
	if MyLinkedList.Head == nil {
		MyLinkedList.AddAtHead(val)
	} else {
		newNode := &Node{Val: val}
		MyLinkedList.Tail.Next = newNode
		MyLinkedList.Tail = newNode
		MyLinkedList.Length++
	}
}

//AddAtIndex a node of value val before the index-th node in the linked list.
//If index equals to the length of linked list, the node will be appended to the end of linked list.
//If index is greater than the length, the node will not be inserted.
func (MyLinkedList *MyLinkedList) AddAtIndex(index int, val int) {
	if index > MyLinkedList.Length || index < 1 {
		return
	}

	if index == MyLinkedList.Length {
		MyLinkedList.AddAtTail(val)
		return
	}

	if index == 1 {
		MyLinkedList.AddAtHead(val)
		return
	}

	current := MyLinkedList.Head
	for 2 < index {
		current = current.Next
		index--
	}

	newNode := &Node{Val: val, Next: current.Next}
	current.Next = newNode
}

//DeleteAtIndex the index-th node in the linked list, if the index is valid.
func (MyLinkedList *MyLinkedList) DeleteAtIndex(index int) {
	if MyLinkedList == nil || index > MyLinkedList.Length || index < 1 {
		return
	}

	current := MyLinkedList.Head
	for index > 1 {
		current = current.Next
		index--
	}

	current = current.Next
	MyLinkedList.Length--
}

//707. Design Linked List
//https://leetcode.com/problems/design-linked-list/
