package relativesortarray

import (
	"testing"

	"github.com/wenndersantos/leetcode-in-go/testutils"
)

func TestRelativeSortArray(t *testing.T) {
	testCases := []struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			arr1: []int{943, 790, 427, 722, 860, 550, 225, 846, 715, 320},
			arr2: []int{943, 715, 427, 790, 860, 722, 225, 320, 846, 550},
			want: []int{943, 715, 427, 790, 860, 722, 225, 320, 846, 550},
		},
		{
			arr1: []int{28, 6, 22, 8, 44, 17},
			arr2: []int{22, 28, 8, 6},
			want: []int{22, 28, 8, 6, 17, 44},
		},
		{
			arr1: []int{2, 21, 43, 38, 0, 42, 33, 7, 24, 13, 12, 27, 12, 24, 5, 23, 29, 48, 30, 31},
			arr2: []int{2, 42, 38, 0, 43, 21},
			want: []int{2, 42, 38, 0, 43, 21, 5, 7, 12, 12, 13, 23, 24, 24, 27, 29, 30, 31, 33, 48},
		},
	}

	for _, tc := range testCases {
		got := RelativeSortArray(tc.arr1, tc.arr2)

		if !testutils.IntSliceAreEqual(got, tc.want) {
			t.Errorf("\narr1: %v\narr2: %v\ngot: %v, want: %v", tc.arr1, tc.arr2, got, tc.want)
		}
	}
}
