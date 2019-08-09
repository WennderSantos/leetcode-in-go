package pairsofsongswithtotaldurationsdivisibleby60

//NumPairsDivisibleBy60 executes in
//Time: O(n) where n is the length of time
//Memory: O(1) fixed size [60]int
func NumPairsDivisibleBy60(time []int) int {
	result, target := 0, 60
	memory := [60]int{}

	for _, songDuration := range time {
		d := songDuration % target
		if d == 0 {
			result += memory[d]
		} else {
			result += memory[target-d]
		}
		memory[d]++
	}

	return result
}

//1010. Pairs of Songs With Total Durations Divisible by 60
//https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/
