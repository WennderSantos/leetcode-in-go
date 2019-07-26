package addtwonumbers

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   &ListNode{0, nil},
			l2:   &ListNode{0, nil},
			want: &ListNode{0, nil},
		},
		{
			l1:   &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			l2:   &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			want: &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
		},
		{
			l1:   &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			l2:   &ListNode{1, &ListNode{2, nil}},
			want: &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		},
		{
			l1:   &ListNode{1, &ListNode{6, &ListNode{4, nil}}},
			l2:   &ListNode{1, &ListNode{7, &ListNode{3, nil}}},
			want: &ListNode{2, &ListNode{3, &ListNode{8, nil}}},
		},
		{
			l1:   &ListNode{1, &ListNode{9, &ListNode{5, nil}}},
			l2:   &ListNode{1, &ListNode{2, nil}},
			want: &ListNode{2, &ListNode{1, &ListNode{6, nil}}},
		},
	}

	for _, tc := range testCases {
		got := AddTwoNumbers(tc.l1, tc.l2)

		for got != nil || tc.want != nil {
			if got == nil || tc.want == nil || got.Val != tc.want.Val {
				t.Errorf("l1: %v\nl2: %v\ngot %v, want %v", tc.l1, tc.l2, got, tc.want)
				break
			}

			got = got.Next
			tc.want = tc.want.Next
		}
	}
}
