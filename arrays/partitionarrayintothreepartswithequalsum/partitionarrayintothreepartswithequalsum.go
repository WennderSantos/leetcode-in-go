package partitionarrayintothreepartswithequalsum

//CanThreePartsEqualSum executes in
//Time: O(n) where n is de length of A
//Memory: O(1)
func CanThreePartsEqualSum(A []int) bool {
	if len(A) < 3 {
		return false
	}

	totalSum := 0
	for _, val := range A {
		totalSum += val
	}

	if totalSum%3 != 0 {
		return false
	}

	partitionSum := totalSum / 3
	numberOfPartitions := 3
	auxSum := 0
	for _, val := range A {
		auxSum += val
		if auxSum == partitionSum {
			auxSum = 0
			numberOfPartitions--
		}

		//just neest to find 2 partitions
		if numberOfPartitions == 1 {
			return true
		}
	}

	return false
}

//1013. Partition Array Into Three Parts With Equal Sum
//https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/
