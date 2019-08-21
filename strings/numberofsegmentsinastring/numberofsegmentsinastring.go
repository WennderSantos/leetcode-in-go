package numberofsegmentsinastring

//CountSegments executes in
//Time: O(n) - where n is the length of s
//Memory: O(1)
func CountSegments(s string) int {
	segments := 0

	for i, val := range s {
		if val != ' ' && (i == 0 || s[i-1] == ' ') {
			segments++
		}
	}
	return segments
}

//434. Number of Segments in a String
//https://leetcode.com/problems/number-of-segments-in-a-string/
