package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			arr:    []int{},
			target: 1,
			want:   0,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 0,
			want:   0,
		},
	}

	for _, tc := range testCases {
		got := SearchInsert(tc.arr, tc.target)

		if got != tc.want {
			t.Errorf("arr: %v, target: %d - got: %d, want: %d", tc.arr, tc.target, got, tc.want)
		}
	}
}
