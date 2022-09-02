/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	cost, profit := 0x7fffffff, 0
	for _, price := range prices {
		if cost > price {
			cost = price
		}
		tmpProfit := price - cost
		if tmpProfit > profit {
			profit = tmpProfit
		}
	}
	return profit
}

// @lc code=end
