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

type ListNode struct {
	Val  int
	Next *ListNode
}

// Reverse a linked list
// https://leetcode.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	// approach 1: iterate through the ll, swap next references to current and vice versa
	// time = O(n) best we can do, space = constant (best)

	// insight: we need to keep track of the current node, the next AND the previous
	// to point "Next" on the current appropriately, so we should actually start at head + 1
	// and keep the actual next in a temp variable

	// edge case: nil head node
	if head == nil {
		return nil
	}

	// edge case: solo node
	if head.Next == nil {
		return head
	}

	prev := head
	current := head.Next
	head.Next = nil

	for current != nil {
		next := current.Next
		current.Next, prev, current = prev, current, next // just go things
	}

	return prev
}

// Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	// key points:
	// detect a "cycle" in the LL where cycle means you can keep following next and arrive at the same node
	// pos is used to (internally) denote the index of the node at the tail end of the cycle (or -1 if no cycle exists) -> tbh doesn't seem helpful if it isn't passed as a param though
	// return true/false of whether the cycle exists

	// can we use multiple pointers?

	// are the values of the nodes unqiue? doesn't seem like we can assume that

	// can I mutate the values in the linked list?
	// even if we can, this doesn't really help us

	return hasCycleThree(head)
}

// hasCycleThree uses fast and slow pointers to find the cycle
// the key insight here being that since we only ever truly "finish"
// iterating for non-cyclical lists and if we are in a cycle, the systematically
// faster pointer should always "meet" the slower pointer eventually
// e.g. fast pointer moves at 2 indices, slow pointer at 1
// both are "in the cycle" of length K, if the fast pointer is behind by 1 step
// the pointers will meet on the next iteration. If the fast pointer is 2, then
// the next iteration will put them in the first position (1 step behind) and so forth
func hasCycleThree(head *ListNode) bool {
	// edge case: nil node or solo node
	if head == nil || head.Next == nil {
		return false
	}

	// edge case: solo node with reference to itself
	if head == head.Next {
		return true
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow, fast = slow.Next, fast.Next.Next
	}

	return false
}

// hasCycleTwo uses two pointers to manually check if the same node
// is referenced twice (i.e. there's a cycle) and returns false otherwise
// time limit exceeded
func hasCycleTwo(head *ListNode) bool {
	// slower but no additional space: look for repeat refs
	// time = O(n^2), space = O(1)
	current := head
	for current != nil {
		temp := current.Next
		for temp != nil {
			if temp.Next == current {
				return true
			}
			temp = temp.Next
		}
		current = current.Next
	}

	return false
}

// hasCycleOne iterates and stores stuff
// if any node reference is duplicated return true, if
func hasCycleOne(head *ListNode) bool {
	// we iterate through em all, then we know no cycle exists
	// time = O(n) best we can do, space = O(n) one mem slot for each node

	mem := make(map[*ListNode]bool)

	current := head
	for current != nil {
		if mem[current] {
			return true
		}
		mem[current] = true
		current = current.Next
	}

	return false
}
