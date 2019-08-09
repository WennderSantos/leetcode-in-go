package buddystrings

import "testing"

func TestBuddyStrings(t *testing.T) {
	testCases := []struct {
		str1 string
		str2 string
		want bool
	}{
		{
			str1: "ab",
			str2: "ba",
			want: true,
		},
		{
			str1: "ab",
			str2: "ab",
			want: false,
		},
		{
			str1: "aa",
			str2: "aa",
			want: true,
		},
		{
			str1: "aaaaaaabc",
			str2: "aaaaaaacb",
			want: true,
		},
		{
			str1: "",
			str2: "aa",
			want: false,
		},
		{
			str1: "aa",
			str2: "",
			want: false,
		},
		{
			str1: "abcaa",
			str2: "abcbb",
			want: false,
		},
	}

	for _, tc := range testCases {
		got := BuddyStrings(tc.str1, tc.str2)

		if got != tc.want {
			t.Errorf("str1: %s, str2: %s - got: %t, want: %t", tc.str1, tc.str2, got, tc.want)
		}
	}
}
