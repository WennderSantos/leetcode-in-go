package duplicatezeros

import (
	"testing"

	"github.com/wenndersantos/leetcode-in-go/testutils"
)

func TestDuplicateZeros(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 0, 2, 3, 0, 4, 5, 0},
			want:  []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		got := DuplicateZeros(tc.input)

		if !testutils.IntSliceAreEqual(got, tc.want) {
			t.Errorf("input: %v - got: %v, want: %v", tc.input, got, tc.want)
		}
	}
}
