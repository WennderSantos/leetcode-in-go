package duplicatezeros

//DuplicateZeros executes in
//Time: O(n^2) - the entire array is 0 in worst case
//Memory: O(1) done in place
func DuplicateZeros(arr []int) []int {
	lastIndex := len(arr) - 1

	for i := lastIndex; i >= 0; i-- {
		if arr[i] == 0 && i != lastIndex {
			aux := arr[i+1]
			arr[i+1] = 0
			j := i + 2

			for j <= lastIndex {
				aux, arr[j] = arr[j], aux
				j++
			}
		}
	}

	return arr
}

//1089. Duplicate Zeros
//https://leetcode.com/problems/duplicate-zeros/
