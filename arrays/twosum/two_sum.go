package twosum

//TwoSum executes in
//time: O(n) where N is the length of nums
//memory: O(n) which is the memory used by the hashmap
func TwoSum(nums []int, target int) []int {
	memory := map[int]int{}

	for i, item := range nums {
		if val, ok := memory[item]; ok {
			return []int{val, i}
		}

		memory[target-item] = i
	}
	return []int{}
}

//1. Two Sum
//https://leetcode.com/problems/two-sum/
