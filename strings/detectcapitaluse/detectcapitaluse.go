package detectcapitaluse

//DetectCapitalUse executes  in
//Time: O(n)
//Memory: O(1)
func DetectCapitalUse(word string) bool {
	foundNonCapital := false
	countCapitals := 0

	for i, val := range word {
		if val >= 65 && val <= 90 {
			if i > 0 && foundNonCapital {
				return false
			}

			countCapitals++
		} else {
			if countCapitals > 1 {
				return false
			}

			foundNonCapital = true
		}
	}

	return true
}

//520. Detect Capital
//https://leetcode.com/problems/detect-capital/submissions/
