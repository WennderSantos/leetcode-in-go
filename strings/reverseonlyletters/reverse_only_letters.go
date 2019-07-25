package reverseonlyletters

func isLetter(asciiChar rune) bool {
	return asciiChar >= 'a' && asciiChar <= 'z' || asciiChar >= 'A' && asciiChar <= 'Z'
}

//ReverseOnlyLetters executes in
//Time: O(n)
//Memory: O(1) using two pointers (low and high) approach
func ReverseOnlyLetters(S string) string {
	i, j, str := 0, len(S)-1, []rune(S)

	for i < j {
		if !isLetter(str[i]) {
			i++
			continue
		}

		if !isLetter(str[j]) {
			j--
			continue
		}

		str[i], str[j] = str[j], str[i]
		i++
		j--
	}

	return string(str)
}

//917. Reverse Only Letters
//https://leetcode.com/problems/reverse-only-letters/
