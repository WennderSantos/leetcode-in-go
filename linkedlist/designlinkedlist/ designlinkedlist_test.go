package designlinkedlist

import (
	"testing"
)

type testCase struct {
	list      MyLinkedList
	input     int
	want      *Node
	newLength int
}

func TestGet(t *testing.T) {
	list := Constructor()
	list.Head = &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}
	list.Length = 3

	testCases := []struct {
		input int
		want  int
	}{
		{
			input: 1,
			want:  1,
		},
		{
			input: 0,
			want:  -1,
		},
		{
			input: 1000,
			want:  -1,
		},
		{
			input: 3,
			want:  3,
		},
	}

	for _, tc := range testCases {
		got := list.Get(tc.input)

		if got != tc.want {
			t.Errorf("input: %v - got: %d, want: %d", tc.input, got, tc.want)
		}
	}
}

func TestAddAtHead(t *testing.T) {
	list := Constructor()
	list.Head = &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}
	list.Length = 3

	testCases := []testCase{
		{
			input:     0,
			want:      &Node{Val: 0, Next: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}},
			newLength: 4,
		},
	}

	for _, tc := range testCases {
		tc.list.AddAtHead(tc.input)
		current := tc.list.Head

		if !listsAreEqual(current, tc.list.Head) {
			t.Errorf("input: %v - got %v, want %v", tc.input, tc.list.Head, tc.want)
		}
		current = current.Next
		tc.want = tc.want.Next
	}
}

func TestAddAtTail(t *testing.T) {
	listWithElements := Constructor()
	listWithElements.AddAtHead(1)
	listWithElements.AddAtTail(2)
	listWithElements.AddAtTail(3)

	testCases := []testCase{
		{
			list:      Constructor(),
			input:     1,
			want:      &Node{Val: 1},
			newLength: 1,
		},
		{
			list:      listWithElements,
			input:     4,
			want:      &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4}}}},
			newLength: 4,
		},
	}

	for _, tc := range testCases {
		tc.list.AddAtTail(tc.input)
		current := tc.list.Head

		if !listsAreEqual(current, tc.list.Head) {
			t.Errorf("input: %v - got %v, want %v", tc.input, tc.list.Head, tc.want)
		}
		current = current.Next
		tc.want = tc.want.Next
	}
}

func TestAddAtIndex(t *testing.T) {
	listWithOneElement := Constructor()
	listWithOneElement.AddAtHead(2)

	listWithTwoElements := Constructor()
	listWithTwoElements.AddAtHead(1)
	listWithTwoElements.AddAtTail(2)

	listWitThreeElements := Constructor()
	listWitThreeElements.AddAtHead(1)
	listWitThreeElements.AddAtTail(3)
	listWitThreeElements.AddAtTail(4)

	listWitFourElements := Constructor()
	listWitFourElements.AddAtHead(1)
	listWitFourElements.AddAtTail(2)
	listWitFourElements.AddAtTail(3)
	listWitFourElements.AddAtTail(4)

	testCases := []struct {
		list  MyLinkedList
		index int
		val   int
		want  *Node
	}{
		{
			list:  Constructor(),
			index: 1,
			val:   1,
			want:  Constructor().Head,
		},
		{
			list:  listWithTwoElements,
			index: 2,
			val:   3,
			want:  &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}},
		},
		{
			list:  listWithOneElement,
			index: 1,
			val:   1,
			want:  &Node{Val: 2, Next: &Node{Val: 1}},
		},
		{
			list:  listWitThreeElements,
			index: 2,
			val:   2,
			want:  &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4}}}},
		},
		{
			list:  listWitFourElements,
			index: 1,
			val:   0,
			want:  &Node{Val: 0, Next: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4}}}}},
		},
	}

	for _, tc := range testCases {
		tc.list.AddAtIndex(tc.index, tc.val)

		if !listsAreEqual(tc.list.Head, tc.want) {
			t.Errorf("index: %d, val: %d - got: %v, want: %v", tc.index, tc.val, tc.list.Head, tc.want)
		}
	}
}

func TestDeleteAtIndex(t *testing.T) {
	listWithOneElement := Constructor()
	listWithOneElement.AddAtHead(2)

	listWithTwoElements := Constructor()
	listWithTwoElements.AddAtHead(1)
	listWithTwoElements.AddAtTail(2)

	listWitThreeElements := Constructor()
	listWitThreeElements.AddAtHead(1)
	listWitThreeElements.AddAtTail(2)
	listWitThreeElements.AddAtTail(3)

	testCases := []struct {
		list  MyLinkedList
		index int
		want  *Node
	}{
		{
			list:  listWithOneElement,
			index: 1,
			want:  Constructor().Head,
		},
		{
			list:  listWithTwoElements,
			index: 2,
			want:  &Node{Val: 1},
		},
		{
			list:  listWitThreeElements,
			index: 2,
			want:  &Node{Val: 1, Next: &Node{Val: 3}},
		},
	}

	for _, tc := range testCases {
		tc.list.DeleteAtIndex(tc.index)
		current := tc.list.Head

		if !listsAreEqual(current, tc.list.Head) {
			t.Errorf("list: %v, index: %d - got %v, want %v", tc.list, tc.index, tc.list.Head, tc.want)
		}
	}
}

func listsAreEqual(a *Node, b *Node) bool {
	for a != nil || b != nil {
		if a == nil || b == nil || a.Val != b.Val {
			return false
		}

		a = a.Next
		b = b.Next
	}
	return true
}
