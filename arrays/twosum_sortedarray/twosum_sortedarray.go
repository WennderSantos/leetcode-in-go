package twosumsortedarray

//TwoSumSortedArray executes in
//Time:  O(n) where N is the length  of numbers
//Memory: O(1) -  As the array is sorted, I am using
//the strategy of two pointers (low and high).
//This exercise on leetcode expects the result indexes in base 1
//thats the why of []int{l + 1, h + 1}
func TwoSumSortedArray(numbers []int, target int) []int {
	l, h := 0, len(numbers)-1

	for l < h {
		if (numbers[l] + numbers[h]) > target {
			h--
		} else if (numbers[l] + numbers[h]) < target {
			l++
		} else {
			return []int{l + 1, h + 1}
		}
	}
	return []int{}
}

//167. Two Sum II - Input array is sorted
//https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
