package main

// list of my answers to a bunch of leetcode questions because version-controlled work > not version-controlled work
// also it feels nice to commit my answers to git

// Two Sum
// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	// approach 1: brute force, iterate through every pair, O(n^2) time, constant space
	// additional insight: we can infer the other number in the array based on i and target

	// approach 2: keep track of numbers we see in a fast-access DS (map? slice?)
	// on each element, look for target - i (meaning the other part of the sum exists in the arr)
	// O(n) time (best we can do), O(n) space

	// can we do better than O(n) space? probably not because the two numbers can be on opposite ends of the array
	// e.g. target = 9, [2,3,4,10,0,7]

	mem := map[int]int{}
	for i, v := range nums {
		if j, ok := mem[target-v]; ok {
			return []int{j, i}
		}

		mem[v] = i
	}

	return []int{-1}
}

// Best Time To Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	// approach 1: brute force, try all combinations of buy/sell, return max difference On^2

	// additional insight: we can keep track of our hypothetical "best" profit as we pass

	// approach 2:
	// - two pointers (buy, sell). start buy at index 0 and sell at index 1.
	// - keep iterating on sell ptr, updating hypothetical profit along the way
	// - if at any point sell < buy, make that the new buy price (but dont update profit)
	// because any index after would be best to buy then, but won't necessarily return more profit
	// - stop when sell ptr reaches the end of the array
	// - return 0 if no profit was set, return profit margin calculated from iteration

	buy, sell := 0, 1

	var profit int
	for sell < len(prices) {
		if prices[sell]-prices[buy] > profit {
			profit = prices[sell] - prices[buy]
		}

		if prices[sell] < prices[buy] {
			buy = sell
		}

		sell++
	}

	return profit
}
