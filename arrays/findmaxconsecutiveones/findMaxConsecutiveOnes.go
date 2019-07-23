package findmaxconsecutiveones

//FindMaxConsecutiveOnes executes in
//Time: O(n)
//Memory: O(1)
func FindMaxConsecutiveOnes(nums []int) int {
	result, count := 0, 0

	for _, val := range nums {
		count *= val
		count += val

		if count > result {
			result = count
		}
	}

	return result
}

//485. Max Consecutive Ones
//https://leetcode.com/problems/max-consecutive-ones/
