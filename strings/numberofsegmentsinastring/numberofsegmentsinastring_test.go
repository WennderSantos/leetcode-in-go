package numberofsegmentsinastring

import "testing"

func TestCountSegments(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: "",
			want:  0,
		},
		{
			input: "  ",
			want:  0,
		},
		{
			input: "Hello, my name is John",
			want:  5,
		},
		{
			input: "Hi there!",
			want:  2,
		},
		{
			input: "Of all the gin joints in all the towns in all the world,   ",
			want:  13,
		},
		{
			input: ", , , ,        a, eaefa",
			want:  6,
		},
	}

	for _, tc := range testCases {
		got := CountSegments(tc.input)

		if got != tc.want {
			t.Errorf("input: %s - got: %d, want: %d", tc.input, got, tc.want)
		}
	}
}
