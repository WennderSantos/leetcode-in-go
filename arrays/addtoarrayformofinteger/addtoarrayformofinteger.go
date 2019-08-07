package addtoarrayformofinteger

//AddToArrayForm executes in
//Time:
//Memory:
func AddToArrayForm(A []int, K int) []int {
	result, j, i := [10001]int{}, 0, len(A)-1

	carry := 0
	for i >= 0 || K > 0 || carry == 1 {
		sum := carry
		if i >= 0 {
			sum += A[i]
			i--
		}

		if K != 0 {
			sum += K % 10
			K /= 10
		}

		result[j] = sum % 10
		carry = sum / 10
		j++
	}

	for l, h := 0, j-1; l < h; l, h = l+1, h-1 {
		result[l], result[h] = result[h], result[l]
	}

	return result[:j]
}

//989. Add to Array-Form of Integer
//https://leetcode.com/problems/add-to-array-form-of-integer/
