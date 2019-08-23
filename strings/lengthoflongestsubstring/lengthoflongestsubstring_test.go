package lengthoflongestsubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: "",
			want:  0,
		},
		{
			input: "a",
			want:  1,
		},
		{
			input: "abcabcbb",
			want:  3,
		},
		{
			input: "bbbbb",
			want:  1,
		},
		{
			input: "pwwkew",
			want:  3,
		},
		{
			input: "dvdf",
			want:  3,
		},
		{
			input: "tmmzuxt",
			want:  5,
		},
	}

	for _, tc := range testCases {
		got := LengthOfLongestSubstring(tc.input)

		if got != tc.want {
			t.Errorf("input: %s - got: %d, want: %d", tc.input, got, tc.want)
		}
	}
}
