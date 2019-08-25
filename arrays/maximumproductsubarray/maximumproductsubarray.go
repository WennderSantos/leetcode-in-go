package maximumproductsubarray

//MaxProduct executes in
//Time: O(n) - where n is the length of nums
//Memory: O(1)
func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	result, currMin, currMax := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			currMax, currMin = currMin, currMax
		}

		if nums[i] > currMax*nums[i] {
			currMax = nums[i]
		} else {
			currMax *= nums[i]
		}

		if nums[i] < currMin*nums[i] {
			currMin = nums[i]
		} else {
			currMin *= nums[i]
		}

		if result < currMax {
			result = currMax
		}

	}
	return result
}

//152. Maximum Product Subarray
//https://leetcode.com/problems/maximum-product-subarray/
