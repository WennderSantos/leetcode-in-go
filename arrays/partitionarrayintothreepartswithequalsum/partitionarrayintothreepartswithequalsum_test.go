package partitionarrayintothreepartswithequalsum

import "testing"

func TestCanThreePartsEqualSum(t *testing.T) {
	testCases := []struct {
		input []int
		want  bool
	}{
		{
			input: []int{1, 1},
			want:  false,
		},
		{
			input: []int{1, 1, 3},
			want:  false,
		},
		{
			input: []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
			want:  true,
		},
		{
			input: []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
			want:  false,
		},
		{
			input: []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4},
			want:  true,
		},
	}

	for _, tc := range testCases {
		got := CanThreePartsEqualSum(tc.input)

		if got != tc.want {
			t.Errorf("input: %v - got: %t, want: %t", tc.input, got, tc.want)
		}
	}

}
