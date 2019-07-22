package findcommonchars

//CommonChars executes in
//Time: O(n^2)
//Memory: O(1) - Three fixed size array [26]int
func CommonChars(A []string) []string {
	toCompare := [26]int{}

	for _, c := range A[0] {
		toCompare[c-'a']++
	}

	for i := 1; i < len(A); i++ {
		countCurrentWord := [26]int{}

		for _, c := range A[i] {
			countCurrentWord[c-'a']++
		}

		for j := 0; j < 26; j++ {
			if countCurrentWord[j] < toCompare[j] {
				toCompare[j] = countCurrentWord[j]
			}
		}
	}

	result := make([]string, 100)
	resultIndex := 0

	for i := 0; i < 26; i++ {
		for j := 0; j < toCompare[i]; j++ {
			result[resultIndex] = string('a' + i)
			resultIndex++
		}
	}

	return result[:resultIndex]
}

//1002. Find Common Characters
