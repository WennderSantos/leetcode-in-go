package numberofequivalentdominopairs

import "testing"

func TestNumEquivDominoPairs(t *testing.T) {
	testCases := []struct {
		input [][]int
		want  int
	}{
		{
			input: [][]int{},
			want:  0,
		},
		{
			input: nil,
			want:  0,
		},
		{
			input: [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}},
			want:  1,
		},
		{
			input: [][]int{{9, 8}, {2, 1}, {8, 9}, {1, 2}},
			want:  2,
		},
		{
			input: [][]int{{1, 2}, {2, 1}, {8, 9}, {1, 2}},
			want:  3,
		},
		{
			input: [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}},
			want:  3,
		},
		{
			input: [][]int{{1, 1}, {2, 2}, {1, 1}, {1, 2}, {1, 2}, {1, 1}},
			want:  4,
		},
		{
			input: [][]int{{2, 1}, {1, 2}, {1, 2}, {1, 2}, {2, 1}, {1, 1}, {1, 2}, {2, 2}},
			want:  15,
		},
	}

	for _, tc := range testCases {
		got := NumEquivDominoPairs(tc.input)

		if got != tc.want {
			t.Errorf("input %v - got: %d, want %d", tc.input, got, tc.want)
		}
	}
}
