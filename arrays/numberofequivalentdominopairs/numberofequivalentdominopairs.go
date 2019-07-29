package numberofequivalentdominopairs

//NumEquivDominoPairs executes in
//Time: O(n) wheren n is the length of dominoes
//Memory: O(1) using a fixed size []int
func NumEquivDominoPairs(dominoes [][]int) int {
	if dominoes == nil || len(dominoes) == 0 {
		return 0
	}

	result := 0
	occurrences := [100]int{}

	for _, dominoe := range dominoes {
		key, rotatedKey := dominoe[0]*10+dominoe[1], dominoe[1]*10+dominoe[0]

		if key != rotatedKey {
			result += occurrences[rotatedKey]
		}

		result += occurrences[key]
		occurrences[key]++
	}

	return result
}

//1128. Number of Equivalent Domino Pairs
//https://leetcode.com/problems/number-of-equivalent-domino-pairs/
