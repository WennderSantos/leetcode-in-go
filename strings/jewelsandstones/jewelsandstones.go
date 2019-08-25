package jewelsandstones

//NumJewelsInStones executes in
//Time: O(j+s)
//Memory: O(1) - fixed size array [256] bool
func NumJewelsInStones(J string, S string) int {
	jewels := [256]bool{}
	for _, c := range J {
		jewels[c] = true
	}

	result := 0
	for _, c := range S {
		if jewels[c] {
			result++
		}
	}
	return result
}

//771. Jewels and Stones
//https://leetcode.com/problems/jewels-and-stones/
