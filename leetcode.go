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

func lengthOfLongestSubstring(s string) int {
	// approach one: sliding window of each character in the string
	// start with index 0, "expand the window" with the second pointer as we iterate
	// store letters we currently have in our substring using a set
	// if repeat found, move start/end and clear set
	// else keep incrementing maxLength, only update max if end - start > currentMax
	// time complexity = O(n) we only look at each character once (best)
	// space complexity = O(n) since our set can be the length of the whole string

	// space is also best I think because how else would we know about repeats without storage?
	// maaaaybe we can do something like XOR'ing the char values to see if there are any repeats
	// but even then, we need to know the actual value of the repeat not just a yes/no

	return lengthOfLongestSubstringOne(s)
}

// worth noting how this performs on ASCII, unicode, UTF-8 encoded strings, etc.?
func lengthOfLongestSubstringOne(s string) int {

	// edge case: empty string or 1 character; just return
	if len(s) <= 1 {
		return len(s)
	}

	maxLength := 0
	mem := make(map[rune]bool, len(s))

	// go strings are really just a slice of runes, so there are some nuances here
	// i.e. this for range loop iterates through the string one rune at a time
	// see: https://go.dev/blog/strings

	for start, end := 0, 0; end < len(s); end++ {
		rEnd := rune(s[end])

		// if repeat, slide start and continuously remove values from mem
		if mem[rEnd] {
			var rStart rune
			for mem[rEnd] {
				rStart = rune(s[start])
				mem[rStart] = false
				start++
			}
		}

		// if size of set > maxLength, update it
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}

		// no matter what, add the current value to mem
		mem[rEnd] = true
	}

	return maxLength
}

// Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/
// todo: work in progress, my brain is tired
func characterReplacement(s string, k int) int {
	// brute force: try every combination of substrings w/ replacements -> O(n^2)

	// approach one: sliding window with map of char -> frequency
	// start, end pointers initialized at 0; map[rune]int containing chars and their frequency
	// very close, but I kept getting stuck on the part when replacements > k
	// fixed this based on a forum answer: https://leetcode.com/problems/longest-repeating-character-replacement/discuss/1341352/Go-Sliding-window-with-comments

	// edge cases? empty string return 0, k > len(s) return len(s) since we can just replace em all

	if len(s) == 0 {
		return 0
	}

	if k > len(s) {
		return len(s)
	}

	start, longestStreak, maxLength := 0, 0, 0
	freq := make(map[rune]int) // map char -> # of ocurrences

	for i, val := range s {
		freq[val] += 1

		// our "longest streak" is the max frequency of the current character
		// keeping track of this lets us know how many replacements we'd
		// theoretically make for every *other* character in our substring
		if freq[val] > longestStreak {
			longestStreak = freq[val]
		}

		// if len(substring) - longestStreak > k means we've replaced too many times
		// increment start, decrement the frequency of that char and hope to find a better streak
		// next iteration

		// note: there's no need to loop here though because we only care about windows
		// at least the size of our best so far, so we can just slide and *not* shrink the window
		if i-start+1-longestStreak > k {
			startChar := rune(s[start])
			freq[startChar]--
			start++
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}
