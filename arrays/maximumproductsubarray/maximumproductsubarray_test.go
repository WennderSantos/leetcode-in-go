package maximumproductsubarray

import "testing"

func TestMaxProduct(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{},
			want:  0,
		},
		{
			input: []int{-2},
			want:  -2,
		},
		{
			input: []int{0, 2},
			want:  2,
		},
		{
			input: []int{2, 3, -2, 4},
			want:  6,
		},
		{
			input: []int{-2, 0, -1},
			want:  0,
		},
		{
			input: []int{-2, 3, -4},
			want:  24,
		},
	}

	for _, tc := range testCases {
		got := MaxProduct(tc.input)

		if got != tc.want {
			t.Errorf("input: %v - got: %d, want: %d", tc.input, got, tc.want)
		}
	}
}
