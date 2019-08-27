package maximumsubarray

import "testing"

func TestMaximumSubarray(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{},
			want:  0,
		},
		{
			input: []int{0},
			want:  0,
		},
		{
			input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want:  6,
		},
	}

	for _, tc := range testCases {
		got := MaxSubArray(tc.input)

		if got != tc.want {
			t.Errorf("input: %v - got: %d, want: %d", tc.input, got, tc.want)
		}
	}
}
