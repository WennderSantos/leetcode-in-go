package findcommonchars

import (
	"testing"

	"github.com/wenndersantos/leetcode-in-go/testutils"
)

func TestFindCommonChars(t *testing.T) {
	testCases := []struct {
		input []string
		want  []string
	}{
		{
			input: []string{"bella", "label", "roller"},
			want:  []string{"e", "l", "l"},
		},
		{
			input: []string{"cool", "lock", "cook"},
			want:  []string{"c", "o"},
		},
	}

	for _, tc := range testCases {
		got := CommonChars(tc.input)

		if !testutils.StringSliceAreEqual(got, tc.want) {
			t.Errorf("input: %v - got: %v, want: %v.", tc.input, got, tc.want)
		}
	}
}
