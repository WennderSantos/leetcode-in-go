package besttimetobuyandsellstock

import "testing"

func TestMaxProfit(t *testing.T) {
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
			want:  0,
		},
		{
			input: []int{7, 1, 5, 3, 6, 4},
			want:  5,
		},
	}

	for _, tc := range testCases {
		got := MaxProfit(tc.input)

		if got != tc.want {
			t.Errorf("input: %v - got: %d, want: %d", tc.input, got, tc.want)
		}
	}
}
