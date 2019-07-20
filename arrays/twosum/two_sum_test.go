package twosum

import (
	"testing"

	"github.com/wenndersantos/leetcode-in-go/testutils"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		arr    []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3, 5, 6, 7}, 4, []int{0, 2}},
		{[]int{1, 2, 3, 5, 6, 7}, 15, []int{}},
	}

	for _, testcase := range cases {
		got := TwoSum(testcase.arr, testcase.target)

		if !testutils.IntSliceAreEqual(got, testcase.want) {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", got, testcase.want)
		}
	}
}
