package thirdmaximumnumber

import "testing"

func TestThirdMax(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{},
			want:  0,
		},
		{
			input: []int{1},
			want:  1,
		},
		{
			input: []int{1, 2},
			want:  2,
		},
		{
			input: []int{3, 2, 1},
			want:  1,
		},
		{
			input: []int{2, 2, 3, 1},
			want:  1,
		},
		{
			input: []int{1, 1, 2},
			want:  2,
		},
		{
			input: []int{1, 2, -2147483648},
			want:  -2147483648,
		},
	}

	for _, tc := range testCases {
		got := ThirdMax(tc.input)

		if got != tc.want {
			t.Errorf("input: %v - got: %d, want: %d", tc.input, got, tc.want)
		}
	}
}
