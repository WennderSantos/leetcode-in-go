package addtoarrayformofinteger

import (
	"testing"

	"github.com/wenndersantos/leetcode-in-go/testutils"
)

func TestAddToArrayForm(t *testing.T) {
	testCases := []struct {
		arr  []int
		k    int
		want []int
	}{
		{
			arr:  []int{1, 2, 0, 0},
			k:    34,
			want: []int{1, 2, 3, 4},
		},
		{
			arr:  []int{2, 7, 4},
			k:    181,
			want: []int{4, 5, 5},
		},
		{
			arr:  []int{2, 1, 5},
			k:    806,
			want: []int{1, 0, 2, 1},
		},
		{
			arr:  []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			k:    1,
			want: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			arr:  []int{0},
			k:    0,
			want: []int{0},
		},
		{
			arr:  []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3},
			k:    516,
			want: []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 5, 7, 9},
		},
		{
			arr:  []int{3, 8, 0, 3, 0, 2, 7, 0, 7, 6, 4, 9, 9, 1, 7, 6, 6, 1, 6, 4},
			k:    670,
			want: []int{3, 8, 0, 3, 0, 2, 7, 0, 7, 6, 4, 9, 9, 1, 7, 6, 6, 8, 3, 4},
		},
	}

	for _, tc := range testCases {
		got := AddToArrayForm(tc.arr, tc.k)

		if !testutils.IntSliceAreEqual(got, tc.want) {
			t.Errorf("\narr %v, k %d\ngot:  %v\nwant: %v", tc.arr, tc.k, got, tc.want)
		}
	}
}
