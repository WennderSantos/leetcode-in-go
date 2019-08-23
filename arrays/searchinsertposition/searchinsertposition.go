package searchinsertposition

//SearchInsert executes in
//Time: O(log n) - binary search
//Memory: O(1)
func SearchInsert(nums []int, target int) int {
	l, h := 0, len(nums)-1

	for l <= h {
		mid := (l + h) / 2

		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}

	return l
}

//35. Search Insert Position
//https://leetcode.com/problems/search-insert-position/submissions/
