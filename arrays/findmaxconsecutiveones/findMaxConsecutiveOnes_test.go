package findmaxconsecutiveones

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1, 1, 0, 1, 1, 1},
			want:  3,
		},
		{
			input: []int{0, 0, 0, 0},
			want:  0,
		},
		{
			input: []int{1, 1, 1, 1},
			want:  4,
		},
	}

	for _, tc := range testCases {
		got := FindMaxConsecutiveOnes(tc.input)

		if got != tc.want {
			t.Errorf("input: %v - got %d, want %d", tc.input, got, tc.want)
		}
	}
}
