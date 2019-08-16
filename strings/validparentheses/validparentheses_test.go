package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			input: "",
			want:  true,
		},
		{
			input: "]",
			want:  false,
		},
		{
			input: "){",
			want:  false,
		},
		{
			input: "[])",
			want:  false,
		},
		{
			input: "()",
			want:  true,
		},
		{
			input: "()[]{}",
			want:  true,
		},
		{
			input: "(]",
			want:  false,
		},
		{
			input: "([)]",
			want:  false,
		},
		{
			input: "{[]}",
			want:  true,
		},
	}

	for _, tc := range testCases {
		got := IsValid(tc.input)

		if got != tc.want {
			t.Errorf("input: %s - got: %t, want: %t", tc.input, got, tc.want)
		}
	}
}
