package main

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and
choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction.
If you cannot achieve any profit, return 0.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
Constraints:
1 <= prices.length <= 105
0 <= prices[i] <= 104
*/

func maxProfit(prices []int) int {
	var minp int
	var profit int

	minp = prices[0]
	profit = 0

	for i := 0; i < len(prices)-1; i++ {
		// if price of next day is less than current day adjust minimum
		//Sliding Window of min and max both are moving
		if prices[i+1] <= prices[i] {
			//new min
			minp = minimumTwo(minp, prices[i+1])

		} else {
			profit = maximumTwo(profit, prices[i+1]-minp)
		}

	}

	return profit
}

func minimumTwo(a int, b int) int {

	if a > b {
		return b
	}
	return a
}

func maximumTwo(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
