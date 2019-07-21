package maxoccurringchar

import (
	"testing"
)

func TestMaxOccurringChar(t *testing.T) {
	cases := []struct {
		text string
		want string
	}{
		{"abc", "a"},
		{"aabc", "a"},
		{"aabccoc", "c"},
		{"aaaccc", "a"},
	}

	for _, testcase := range cases {
		got := MaxOccurringChar(testcase.text)

		if got != testcase.want {
			t.Errorf("got: %s, want: %s.", got, testcase.want)
		}
	}
}
