package reverseonlyletters

import "testing"

func TestReverseOnlyLetters(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "ab-cd",
			want:  "dc-ba",
		},
		{
			input: "a-bC-dEf-ghIj",
			want:  "j-Ih-gfE-dCba",
		},
		{
			input: "Test1ng-Leet=code-Q!",
			want:  "Qedo1ct-eeLg=ntse-T!",
		},
	}

	for _, tc := range testCases {
		got := ReverseOnlyLetters(tc.input)

		if got != tc.want {
			t.Errorf("input: %s - got: %s, want: %s.", tc.input, got, tc.want)
		}
	}
}
