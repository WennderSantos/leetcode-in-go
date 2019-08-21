package plusone

import (
	"testing"

	"github.com/wenndersantos/leetcode-in-go/testutils"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{},
			want:  []int{},
		},
		{
			input: []int{1},
			want:  []int{2},
		},
		{
			input: []int{1, 2, 3},
			want:  []int{1, 2, 4},
		},
		{
			input: []int{4, 3, 2, 1},
			want:  []int{4, 3, 2, 2},
		},
		{
			input: []int{1, 2, 9},
			want:  []int{1, 3, 0},
		},
		{
			input: []int{9, 9, 9},
			want:  []int{1, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := PlusOne(tc.input)

		if !testutils.IntSliceAreEqual(got, tc.want) {
			t.Errorf("input: %v - got: %v, want: %v", tc.input, got, tc.want)
		}
	}
}
