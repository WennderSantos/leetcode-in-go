package maxoccurringchar

//MaxOccurringChar executes in
//Time: O(n) where n is the length of the input.
//Memory: O(1) - Always alocating a fixed size hashmap
//doesn't matter length of n
func MaxOccurringChar(text string) string {
	result := ""

	if len(text) == 0 {
		return result
	}

	memory := make(map[rune]int, 62)
	max := 0

	for _, char := range text {
		if _, ok := memory[char]; ok {
			memory[char]++
		} else {
			memory[char] = 1
		}

		if memory[char] > max {
			result = string(char)
			max = memory[char]
		}
	}

	return result
}
