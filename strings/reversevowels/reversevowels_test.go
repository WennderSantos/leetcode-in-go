package reversevowels

import "testing"

func TestReverseVowels(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "",
			want:  "",
		},
		{
			input: "p",
			want:  "p",
		},
		{
			input: "po",
			want:  "po",
		},
		{
			input: "Aa",
			want:  "aA",
		},
		{
			input: "hello",
			want:  "holle",
		},
		{
			input: "leetcode",
			want:  "leotcede",
		},
	}

	for _, tc := range testCases {
		got := ReverseVowels(tc.input)

		if got != tc.want {
			t.Errorf("input: %s - got: %s, want: %s", tc.input, got, tc.want)
		}
	}
}
