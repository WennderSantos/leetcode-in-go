package maxoccurringchar

//MaxOccurringChar executes in
//Time: O(n) where n is the length of the input
//Memory: O(1) - Always alocating [256]int,
//doesn't matter length of n
func MaxOccurringChar(text string) string {
	result := ""

	if len(text) == 0 {
		return result
	}

	memory := [256]int{}
	max := 0

	for _, char := range text {
		memory[char]++

		if memory[char] > max {
			result = string(char)
			max = memory[char]
		}
	}

	return result
}
