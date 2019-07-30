package removeduplicatesfromsortedlist

import (
	"testing"

	ln "github.com/wenndersantos/leetcode-in-go/linkedlist/listnode"
)

func TestDeleteDuplicates(t *testing.T) {
	testCases := []struct {
		input *ln.ListNode
		want  *ln.ListNode
	}{
		{
			input: &ln.ListNode{},
			want:  &ln.ListNode{},
		},
		{
			input: &ln.ListNode{Val: 1, Next: nil},
			want:  &ln.ListNode{Val: 1, Next: nil},
		},
		{
			input: &ln.ListNode{Val: 1, Next: &ln.ListNode{Val: 1, Next: nil}},
			want:  &ln.ListNode{Val: 1, Next: nil},
		},
		{
			input: &ln.ListNode{Val: 1, Next: &ln.ListNode{Val: 2, Next: nil}},
			want:  &ln.ListNode{Val: 1, Next: &ln.ListNode{Val: 2, Next: nil}},
		},
		{
			input: &ln.ListNode{Val: 1, Next: &ln.ListNode{Val: 1, Next: &ln.ListNode{Val: 2, Next: nil}}},
			want:  &ln.ListNode{Val: 1, Next: &ln.ListNode{Val: 2, Next: nil}},
		},
	}

	for _, tc := range testCases {
		got := DeleteDuplicates(tc.input)

		for got != nil || tc.want != nil {
			if got == nil || tc.want == nil || got.Val != tc.want.Val {
				t.Errorf("input: %v - got %v, want %v", tc.input, got, tc.want)
				break
			}

			got = got.Next
			tc.want = tc.want.Next
		}
	}
}
