package validparentheses

//IsValid executes in
//Time: O(n) - where n is the length of the input
//Memory: O(1) - I am using a fixed size stack,
//but would be better to use one with variable
//length which would result in O(n) memory
func IsValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	stack, topIndex := [100]rune{}, -1
	kv := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, val := range s {
		switch val {
		case '(', '{', '[':
			topIndex++
			stack[topIndex] = val
		case ')', '}', ']':
			if topIndex == -1 || stack[topIndex] != kv[val] {
				return false
			}
			topIndex--
		}
	}
	return topIndex == -1
}

//20. Valid Parentheses
//https://leetcode.com/problems/valid-parentheses/
