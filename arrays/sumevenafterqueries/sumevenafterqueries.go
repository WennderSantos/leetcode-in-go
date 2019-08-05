package sumevenafterqueries

//SumEvenAfterQueries executes in
//Time:
//Memory: O(1) - fixed sized result array
func SumEvenAfterQueries(A []int, queries [][]int) []int {
	sum := 0

	for _, val := range A {
		if val%2 == 0 {
			sum += val
		}
	}

	result := [10001]int{}

	for i, querie := range queries {
		idx, val := querie[1], querie[0]
		if A[idx]%2 == 0 {
			sum -= A[idx]
		}

		A[idx] += val

		if A[idx]%2 == 0 {
			sum += A[idx]
		}

		result[i] = sum
	}

	return result[:len(queries)]
}

//985. Sum of Even Numbers After Queries
//https://leetcode.com/problems/sum-of-even-numbers-after-queries/
