package maximumsubarray

//MaxSubArray executes in
//Time: O(n) - where n is the length of nums
//Memory: O(1)
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result, currMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > currMax+nums[i] {
			currMax = nums[i]
		} else {
			currMax += nums[i]
		}

		if currMax > result {
			result = currMax
		}
	}

	return result
}

//53. Maximum Subarray
//https://leetcode.com/problems/maximum-subarray/
