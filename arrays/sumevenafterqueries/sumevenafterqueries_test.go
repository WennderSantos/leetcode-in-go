package sumevenafterqueries

import (
	"testing"

	"github.com/wenndersantos/leetcode-in-go/testutils"
)

func TestSumEvenAfterQueries(t *testing.T) {
	testCases := []struct {
		arr     []int
		queries [][]int
		want    []int
	}{
		{
			arr:     []int{1, 2, 3, 4},
			queries: [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}},
			want:    []int{8, 6, 2, 4},
		},
	}

	for _, tc := range testCases {
		got := SumEvenAfterQueries(tc.arr, tc.queries)

		if !testutils.IntSliceAreEqual(got, tc.want) {
			t.Errorf("\narr: %v\nqueries: %v\ngot: %v\nwant: %v", tc.arr, tc.queries, got, tc.want)
		}
	}
}
