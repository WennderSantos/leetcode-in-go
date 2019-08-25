package singlerowkeyboard

import "testing"

func TestCalculateTime(t *testing.T) {
	testCases := []struct {
		keyboard string
		word     string
		want     int
	}{
		{
			keyboard: "abcdefghijklmnopqrstuvwxyz",
			word:     "cba",
			want:     4,
		},
		{
			keyboard: "pqrstuvwxyzabcdefghijklmno",
			word:     "leetcode",
			want:     73,
		},
	}

	for _, tc := range testCases {
		got := CalculateTime(tc.keyboard, tc.word)

		if got != tc.want {
			t.Errorf("keyboard: %s, word: %s - got: %d, want: %d", tc.keyboard, tc.word, got, tc.want)
		}
	}
}
