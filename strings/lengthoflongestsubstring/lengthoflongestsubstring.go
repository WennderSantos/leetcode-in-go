package lengthoflongestsubstring

//LengthOfLongestSubstring executes in
//Time: O(n) - where n is the length of s
//Memory: O(1) - using a fixed size [256]int
func LengthOfLongestSubstring(s string) int {
	result, l, r, buffer := 0, 0, 0, [256]int{}

	for r < len(s) {
		if buffer[s[r]] > 0 {
			if l < buffer[s[r]] {
				l = buffer[s[r]]
			}
		}

		if r-l+1 > result {
			result = r - l + 1
		}

		buffer[s[r]] = r + 1
		r++
	}
	return result
}

//3. Longest Substring Without Repeating Characters
//https://leetcode.com/problems/longest-substring-without-repeating-characters/
