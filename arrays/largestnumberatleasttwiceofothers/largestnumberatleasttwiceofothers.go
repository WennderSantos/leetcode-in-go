package largestnumberatleasttwiceofothers

//DominantIndex executes in
//Time: O(n) where n is the length of nums
//Memory: O(1)
func DominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	i, biggest, secondBiggest := 1, 0, 0

	for i < len(nums) {
		if nums[biggest] < nums[i] {
			secondBiggest = biggest
			biggest = i
		} else if secondBiggest == biggest || nums[secondBiggest] < nums[i] {
			secondBiggest = i
		}
		i++
	}

	if nums[biggest] >= nums[secondBiggest]*2 {
		return biggest
	}

	return -1
}

//747. Largest Number At Least Twice of Others
//https://leetcode.com/problems/largest-number-at-least-twice-of-others/
