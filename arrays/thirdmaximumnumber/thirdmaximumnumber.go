package thirdmaximumnumber

import "math"

//ThirdMax executes in
//Time: O(n) where n is the length of nums
//Memory: O(1)
//This could also be done using fixed size heap in time O(n log n)
func ThirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	first, second, third := math.MinInt32, math.MinInt32, math.MinInt32
	foundMinIntInArray, count := 0, 0

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if val == first || val == second || val == third {
			if val == math.MinInt32 {
				foundMinIntInArray = 1
			}
			continue
		}

		if val > first {
			third, second, first = second, first, val
		} else if val > second {
			third, second = second, val
		} else if val > third {
			third = val
		}

		count++
	}

	if count+foundMinIntInArray > 2 {
		return third
	}
	return first //there is no third maximum, so return maximum value
}

//414. Third Maximum Number
//https://leetcode.com/problems/third-maximum-number/
