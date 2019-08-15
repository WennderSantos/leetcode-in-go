package reversewordsinastringiii

import "testing"

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "Let's take LeetCode contest",
			want:  "s'teL ekat edoCteeL tsetnoc",
		},
		{
			input: "p",
			want:  "p",
		},
		{
			input: "",
			want:  "",
		},
	}

	for _, tc := range testCases {
		got := ReverseWords(tc.input)

		if got != tc.want {
			t.Errorf("input: %s\ngot: %s\nwant: %s", tc.input, got, tc.want)
		}
	}
}
