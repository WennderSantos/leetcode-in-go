package buddystrings

//BuddyStrings executes in
//Time: O(n) where n where is the length of A
//Memory: O(1)
func BuddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	memory := [26]int{}
	firstIdx, secondIdx := -1, -1

	if A == B {
		for _, val := range A {
			memory[val-'a']++

			if memory[val-'a'] > 1 {
				return true
			}
		}
	} else {
		for i := 0; i < len(A); i++ {
			if A[i] != B[i] {
				if firstIdx == -1 {
					firstIdx = i
				} else if secondIdx == -1 {
					secondIdx = i
				} else { // more than 1 swap
					return false
				}
			}
		}

	}
	return secondIdx != -1 && A[firstIdx] == B[secondIdx] && A[secondIdx] == B[firstIdx]
}

//859. Buddy Strings
//https://leetcode.com/problems/buddy-strings/
