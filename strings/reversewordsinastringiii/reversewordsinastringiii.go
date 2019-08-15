package reversewordsinastringiii

//ReverseWords executes in
//Time: O(n * m) where N is the length of string and m is the length of the biggest word
//Memory: O(n) wheren n is result. (should I consider?)
func ReverseWords(s string) string {
	if len(s) <= 1 {
		return s
	}

	i, j, result := 0, 0, []byte(s)

	for j <= len(s) {
		if j == len(s) && i < j || s[j] == ' ' {
			aux := j - 1
			for aux > i {
				result[i], result[aux] = s[aux], s[i]
				aux--
				i++
			}
			i = j + 1
		}
		j++
	}

	return string(result)
}

//557. Reverse Words in a String III
//https://leetcode.com/problems/reverse-words-in-a-string-iii/
