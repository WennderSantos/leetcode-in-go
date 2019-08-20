package reversevowels

import "strings"

func isVowel(c rune) bool {
	vowels := "aeiouAEIOU"

	return strings.ContainsRune(vowels, c)
}

//ReverseVowels executes in
//Time: O(n) - where n is the length of s. In the worst case, there is no vowel
//Memory: O(1)
func ReverseVowels(s string) string {
	i, j, result := 0, len(s)-1, []rune(s)

	for i < j {
		for i < j && !isVowel(result[i]) {
			i++
		}
		for i < j && !isVowel(result[j]) {
			j--
		}
		result[j], result[i] = result[i], result[j]
		j--
		i++
	}

	return string(result)
}

//345. Reverse Vowels of a String
//https://leetcode.com/problems/reverse-vowels-of-a-string/
