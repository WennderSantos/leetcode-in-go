package pairsofsongswithtotaldurationsdivisibleby60

import "testing"

func TestNumPairsDivisibleBy60(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{},
			want:  0,
		},
		{
			input: []int{30, 20, 150, 100, 40},
			want:  3,
		},
		{
			input: []int{60, 60, 60},
			want:  3,
		},
	}

	for _, tc := range testCases {
		got := NumPairsDivisibleBy60(tc.input)

		if got != tc.want {
			t.Errorf("input: %v - got: %d, want: %d", tc.input, got, tc.want)
		}
	}
}
