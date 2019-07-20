package twosumsortedarray

import (
	"testing"

	"github.com/wenndersantos/leetcode-in-go/testutils"
)

func TestTwoSumSortedArray(t *testing.T) {
	cases := []struct {
		arr    []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3, 5, 6, 7}, 4, []int{1, 3}},
		{[]int{1, 2, 3, 5, 6, 7, 20}, 15, []int{}},
	}

	for _, testcase := range cases {
		got := TwoSumSortedArray(testcase.arr, testcase.target)

		if !testutils.IntSliceAreEqual(got, testcase.want) {
			t.Errorf("input: %d - got: %d, want: %d.", testcase.arr, got, testcase.want)
		}
	}
}
