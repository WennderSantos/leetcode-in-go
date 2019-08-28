package besttimetobuyandsellstock

//MaxProfit executes in
//Time: O(n) where n is the length of prices
//Memory: O(1)
func MaxProfit(prices []int) int {
	result, buyAt := 0, 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[buyAt] {
			buyAt = i
		} else if (prices[i] - prices[buyAt]) > result {
			result = prices[i] - prices[buyAt]
		}
	}
	return result
}

//121. Best Time to Buy and Sell Stock
//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
