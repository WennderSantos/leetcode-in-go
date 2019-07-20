package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, h := 0, len(numbers)-1

	for l < h {
		if (numbers[l] + numbers[h]) > target {
			h--
		} else if (numbers[l] + numbers[h]) < target {
			l++
		} else {
			return []int{l + 1, h + 1}
		}
	}
	return []int{}
}

func main() {
	param := []int{1, 2, 3, 4, 5, 6, 7, 8}
	target := 9
	result := twoSum(param, target)

	fmt.Println(result) //[1 8]
}

//167. Two Sum II - Input array is sorted
