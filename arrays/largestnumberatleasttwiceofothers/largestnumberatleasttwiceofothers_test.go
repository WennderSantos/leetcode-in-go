package largestnumberatleasttwiceofothers

import "testing"

func TestDominantIndex(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1},
			want:  0,
		},
		{
			input: []int{1, 0},
			want:  0,
		},
		{
			input: []int{0, 0, 0, 1},
			want:  3,
		},
		{
			input: []int{3, 6, 1, 0},
			want:  1,
		},
		{
			input: []int{1, 2, 3, 4},
			want:  -1,
		},
	}

	for _, tc := range testCases {
		got := DominantIndex(tc.input)

		if got != tc.want {
			t.Errorf("input: %v - got: %d, want: %d", tc.input, got, tc.want)
		}
	}
}
