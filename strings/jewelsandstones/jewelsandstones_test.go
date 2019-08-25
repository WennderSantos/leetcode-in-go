package jewelsandstones

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	testCases := []struct {
		jewels   string
		myStones string
		want     int
	}{
		{
			jewels:   "aA",
			myStones: "aAAbbbb",
			want:     3,
		},
		{
			jewels:   "z",
			myStones: "ZZ",
			want:     0,
		},
	}

	for _, tc := range testCases {
		got := NumJewelsInStones(tc.jewels, tc.myStones)

		if got != tc.want {
			t.Errorf("jewels: %s, myStones: %s - got: %d, want: %d", tc.jewels, tc.myStones, got, tc.want)
		}
	}
}
