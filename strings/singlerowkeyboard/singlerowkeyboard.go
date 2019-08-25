package singlerowkeyboard

//CalculateTime executes in
//Time: O(n) where n is the length of word. Assuming keyboard will always have a fixed length
//Memory:O(1) - Assuming keyboard will always have a fixed length
func CalculateTime(keyboard string, word string) int {
	k := make(map[rune]int, len(keyboard))
	for i, c := range keyboard {
		k[c] = i
	}

	result, lastIndex := 0, 0
	for _, c := range word {
		diff := lastIndex - k[c]
		if diff < 0 {
			diff *= -1
		}
		result += diff
		lastIndex = k[c]
	}
	return result
}

//1165. Single-Row Keyboard
//https://leetcode.com/problems/single-row-keyboard/
