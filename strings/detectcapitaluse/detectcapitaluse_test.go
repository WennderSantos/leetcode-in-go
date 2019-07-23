package detectcapitaluse

import "testing"

func TestDetectCapitalUse(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			input: "foobar",
			want:  true,
		},
		{
			input: "FOOBAR",
			want:  true,
		},
		{
			input: "Foobar",
			want:  true,
		},
		{
			input: "foobaR",
			want:  false,
		},
		{
			input: "fooBar",
			want:  false,
		},
	}

	for _, tc := range testCases {
		got := DetectCapitalUse(tc.input)

		if got != tc.want {
			t.Errorf("input: %v - got %t, want %t", tc.input, got, tc.want)
		}
	}
}
