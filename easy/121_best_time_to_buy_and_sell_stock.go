//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

package easy

// @lc code=start
func maxProfit(prices []int) int {
	// pointer to the day we buy the stock
	// and the current day we check
	buyingDay, currentDay := 0, 0

	// no profit yet
	profit := 0

	// loop until no more day to check
	for currentDay < len(prices) {
		// buy stock if stock is cheaper
		if prices[currentDay] < prices[buyingDay] {
			buyingDay = currentDay
		} else {
			// else sell our stock
			profit = max(profit, prices[currentDay]-prices[buyingDay])
		}

		currentDay++
	}

	return profit
}

// @lc code=end
