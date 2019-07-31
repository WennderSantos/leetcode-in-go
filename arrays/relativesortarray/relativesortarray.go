package relativesortarray

import (
	"sort"
)

//RelativeSortArray executes in
//Time: O(n^2)
//Memory: O(1) - done in place
func RelativeSortArray(arr1 []int, arr2 []int) []int {
	if len(arr1) == 0 || len(arr2) == 0 {
		return arr1
	}

	i, j := 0, 0

	for _, val := range arr2 {
		for i < len(arr1) && arr1[i] == val {
			i++
		}

		j = i + 1

		for j < len(arr1) {
			if arr1[j] == val {
				arr1[j], arr1[i] = arr1[i], arr1[j]
				i++
			}
			j++
		}
	}

	if i < len(arr1) {
		sort.Sort(sort.IntSlice(arr1[i:]))
	}

	return arr1
}

//RelativeSortArray2 executes in
//Time: O(n*Logn) - sort
//Memory: O(n) - where n is the length of arr1
func RelativeSortArray2(arr1 []int, arr2 []int) []int {
	if len(arr1) == 0 || len(arr2) == 0 {
		return arr1
	}

	memory := map[int]int{}
	for _, val := range arr1 {
		if _, ok := memory[val]; ok {
			memory[val]++
		} else {
			memory[val] = 1
		}
	}

	i := 0
	result := [1000]int{}
	for _, val := range arr2 {
		if ocurrences, ok := memory[val]; ok {
			for ocurrences > 0 {
				result[i] = val
				i++
				ocurrences--
			}

			delete(memory, val)
		}
	}

	j := i
	for key, ocurrences := range memory {
		for ocurrences > 0 {
			result[i] = key
			i++
			ocurrences--
		}

		delete(memory, key)
	}

	if j < len(arr1) {
		sort.Sort(sort.IntSlice(result[j:i]))
	}

	return result[:i]
}

//1122. Relative Sort Array
//https://leetcode.com/problems/relative-sort-array/
