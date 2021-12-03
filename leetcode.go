package main

import (
	"container/heap"
	"sort"
	"strconv"
	"strings"
)

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

// Sum of Two Integers
// https://leetcode.com/problems/sum-of-two-integers/
// note to self: I think these are good to know and kinda interesting, but I doubt I'll actually be asked these directly
// mainly because they oftentimes read like brain teasers and say way more about prep/niche knowledge
// than actual problem solving skills
func getSum(a int, b int) int {
	// approach one: naive bitwise addition
	// iterate through a and b bit by bit (right most bit first)
	// XOR the bits to get the next bit for the sum (including carry)
	// AND the bits to figure out the next carry bit

	// close but no dice!

	// the "elegant" solution provided (broken down for me to learn)
	for b != 0 {
		// a = XORing a and b effectively gives us the bit we need when adding each
		// now the problem is reduced to finding the carry...
		// ANDing a and b and shifting left does just that
		// the case where 1 & 1 is the only time we'd have a carry
		// we itearte through our values for a and b in sync because the bits line up
		// as we XOR "add without carry" and compute the AND for the next iteration "find the carry"
		a, b = a^b, (a&b)<<1
	}

	return a
}

// merge K sorted lists
// https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}

	// approach 1: use a min heap to insert the data into one heap then extract em
	// n = total elements across all linked lists to be merged
	// time complexity = heapify (nlogn) + extract all (nlogn)
	// space complexity = heap itself (n) + resulting linked list (n)

	return mergeKListsOne(lists)

	// additional insight: we can also advantage of the fact that each sub list is sorted

	// approach 2: use two pointers to merge two lists and repeat for each subsequent list
	// takes advantage of the fact that each list are sorted in ascending order
	// n = total elements across all linked lists to be merged
	// time complexity = O(n^2) we have to compare nodes to the intermediate merged list
	// space complexity = (not including resulting linked list) constant space

	// (from solution post) also can apply approach 2 in a divide&conquer manner, which I think
	// is a pretty cool approach actually
}

func mergeKListsOne(lists []*ListNode) *ListNode {
	// heapify data
	heap := MinHeap{}
	for _, list := range lists {
		node := list
		for node != nil {
			heap.Insert(node.Val)
			node = node.Next
		}
	}

	if len(heap.arr) == 0 {
		return nil
	}

	// extract min & build a new singly linked list
	head := &ListNode{
		Val:  heap.Extract(),
		Next: nil,
	}

	current := head
	for len(heap.arr) > 0 {
		min := heap.Extract()
		current.Next = &ListNode{
			Val:  min,
			Next: nil,
		}
		current = current.Next
	}

	return head
}

// Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	// approach one: use priority queue where more frequent == higher priority; store then pop k times
	// time complexity: get frequencies = O(n), heapify = O(nlogn), extract top k = O(klogn) -> O(nlogn)
	// space complexity: map/heap space = number of unique elements so O(n) worst case
	return topKFrequentApproachOne(nums, k)

	// can this get better time/space wise?

	// only a slight optimization, sorting is the bottleneck, is it possible to do this without s?

	// approach two: use a better/tailored sorting algorithm that isn't nlogn (bucket sort?)
	// first pass: run through nums, store frequency counts in map (int->int)
	// second pass: run through map entires, store values in buckets based on frequencies
	// third pass: run through buckets (iterate to highest frequencies first), add elements until
	// k limit is hit
	// time complexity: 3 passes at most n elements in each -> O(3n) = O(n)
	// space complexity: frequencies map = # unique elements in nums, buckets = same
	// both would be n worst case = (n)

	// return topKFrequentApproachOne(nums, k) <- slightly faster, but way worse space wise according
	// to leetcode stats lol
}

func topKFrequentApproachOne(nums []int, k int) []int {
	// iterate through nums to get their frequencies
	freqs := make(map[int]int)
	for _, v := range nums {
		freqs[v]++
	}

	// add each element to the pq
	pq := make(PriorityQueue, len(freqs))
	i := 0
	for k, v := range freqs {
		pq[i] = &Item{
			value:    strconv.Itoa(k),
			priority: v,
			index:    i,
		}
		i++
	}

	heap.Init(&pq)

	// while k > 0: extract max (most frequent element) from the pq
	result := []int{}
	for k > 0 {
		item := heap.Pop(&pq).(*Item)

		itemVal, err := strconv.Atoi(item.value)
		if err != nil {
			return nil
		}

		result = append(result, itemVal)
		k--
	}

	return result
}

func topKFrequentApproachTwo(nums []int, k int) []int {
	// iterate through nums to get their frequencies number -> frequency
	freqs := make(map[int]int)
	for _, v := range nums {
		freqs[v]++
	}

	// put numbers into buckets where index = frequency of that value
	// since frequency of any number is at most len(nums), we init with that many buckets
	// but to get around zero actually being a valid element (with some high frequency), we bucket keys as strings
	// and convert back when we're appending to our result
	l := len(nums)
	buckets := make([]string, l)
	for key, val := range freqs {
		// we index on frequency so that the order of our elements imply the most frequent nums
		// and to avoid iterating through the buckets backwards, we index by len(nums) - frequency - 1
		buckets[l-1-val] = strconv.Itoa(key)
	}

	// now just go through our buckets and add every non-empty value to our result until we reach k
	i := 0
	result := make([]int, k)
	for k > 0 {
		if buckets[i] != "" {
			top, err := strconv.Atoi(buckets[i])
			if err != nil {
				panic(err)
			}
			result = append(result, top)
		}
		k--
		i++
	}

	return result
}

func topKFrequentApproachTwo(nums []int, k int) []int {
	// edge case: if nums has only 1 element, we know the answer
	if len(nums) < 2 {
		return nums
	}

	// iterate through nums to get their frequencies number -> frequency
	freqs := make(map[int]int)
	for _, v := range nums {
		freqs[v]++
	}

	// put numbers into buckets where index = frequency of that value
	// since frequency of any number is at most len(nums), we init with that many buckets
	// but to get around zero actually being a valid element (with some high frequency), we bucket keys as strings
	// and convert back when we're appending to our result
	l := len(nums)
	buckets := make([]string, l+1)

	for key, val := range freqs {
		// we index on frequency so that the order of our elements imply the most frequent nums
		// and to avoid iterating through the buckets backwards, we index by len(nums) - frequency - 1
		currentVal := buckets[l-(val-1)]

		if currentVal != "" {
			// concatenate them if another value has the same frequency
			buckets[l-(val-1)] = buckets[l-(val-1)] + "," + strconv.Itoa(key)
		} else {
			buckets[l-(val-1)] = strconv.Itoa(key)
		}

	}

	// now just go through our buckets and add every non-empty value to our result until we reach k

	i := 0
	result := []int{}

	for k > 0 {
		if buckets[i] != "" {
			values := strings.Split(buckets[i], ",")

			for _, v := range values {
				if k == 0 {
					break
				}

				top, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				result = append(result, top)
				k--
			}
		}
		i++
	}

	return result
}

func containsDuplicate(nums []int) bool {
	// approach one: store nums in a set, break and return if we find a duplicate
	// return containsDuplicateOne(nums)
	// time = O(n) best we can do since we need to see every element to know
	// space = O(n) worst case nums is an array of all unique elements (set would be n)

	// can we do better space wise?
	// additional insight: keep a running value that tells us if some number was seen

	// approach two: if our nums was sorted, we could return as soon as we get a num 2x
	// time = O(nlogn), space = O(1)
	return containsDuplicateTwo(nums)
}

func containsDuplicateOne(nums []int) bool {
	set := make(map[int]bool)
	for _, v := range nums {
		if set[v] {
			return true // duplicate num found
		}
		set[v] = true
	}

	return false
}

func containsDuplicateTwo(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	// sort nums
	sort.Ints(nums)

	// look for a streak of nums
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		current := nums[i]
		if prev == current {
			return true
		}

		prev = current
	}

	return false
}

// Product Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/submissions/
func productExceptSelf(nums []int) []int {
	// approach one (realistic): first pass multiple all of them, second pass divide by i
	// stupid constraint to not use '/' though ...

	// approach two (brute force): for each element go through every other element to get
	// answer[i]. time = O(n^2), space = O(1)

	// can I update all "ongoing" products without visiting/comparing an index more than
	// a few times?

	// approach three: "windshield wiper" -> iterate forwards and backwards, until the mid
	// on each swipe, we can update the the "end" index with what we've computed
	// time = still O(n^2) but slightly better, space = O(1)

	// we don't need to compute products all over again though?
	// I know everything before a number and everything after it in the array

	// approach four: (aha moment) windshield wiper actually doesn't need to converge
	// if we store every running product iterating from left -> right
	// we get partial answers for (0, i]
	// if we do the same backwards, our running product would be (i, n -1)
	// the product of both sub answers == the actual product
	return productExceptSelfApproachFour(nums)
}

func productExceptSelfApproachFour(nums []int) []int {
	l := len(nums)
	answer := make([]int, l)

	// iterate from left to right
	p1 := 1
	for i := 0; i < l; i++ {
		answer[i] = p1
		p1 *= nums[i]
	}

	// iterate from right to left
	p2 := 1
	for j := l - 1; j >= 0; j-- {
		answer[j] *= p2
		p2 *= nums[j]
	}

	return answer
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	// approach one: classic algorithm to find the height of a binary tree (recursive)
	// base case: root == nil -> return 0
	// "process the current node": if left != nil -> go deeper, (+ same for the right)
	// return max(left, right) + 1

	// example
	// tree = [3,9,20,null,null,15,7]

	// "height" technically counts the number of edges and we want to count nodes
	// so let's add 1 to include the root node
	return maxDepthHelper(root, 1)
}

func maxDepthHelper(root *TreeNode, depth int) int {
	// base case
	if root == nil {
		return 0
	}

	// "process" the current node
	var heightL, heightR int
	if root.Left != nil {
		heightL = maxDepthHelper(root.Left, depth) // double check on the depth val here
	}

	if root.Right != nil {
		heightR = maxDepthHelper(root.Right, depth) // double check on the depth val here
	}

	return max(heightL, heightR) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Same Tree
// https://leetcode.com/problems/same-tree/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// same if they are "structurally" identical and the nodes have the same value

	// approach one: dfs both trees side by side, if we find a hiccup return
	// base case: if p == nil && q == nil -> return true
	// structural check: if only one is nil -> return false (they should be the same)
	// value check: if p.Val != q.Val -> return false
	// return (stack frame build up): return dfs(p.left, q.left) && dfs(p.right, q.right)
	return isSameTreeDFS(p, q)

	// approach two: do BFS instead
	// return isSameTreeBFS(p, q)
}

func isSameTreeDFS(p *TreeNode, q *TreeNode) bool {
	// base case
	if p == nil && q == nil {
		return true
	}

	// structural check
	if p == nil || q == nil {
		return false
	}

	// node value check
	if p.Val != q.Val {
		return false
	}

	return isSameTreeDFS(p.Left, q.Left) && isSameTreeDFS(p.Right, q.Right)
}

func isSameTreeBFS(p *TreeNode, q *TreeNode) bool {
	// base case
	if p == nil && q == nil {
		return true
	}

	// structural check
	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}

	// keep track of tree nodes to read using two separate queues
	queue1, queue2 := make([]*TreeNode, 0), make([]*TreeNode, 0)
	queue1, queue2 = append(queue1, p), append(queue2, q)

	for len(queue1) > 0 && len(queue2) > 0 {
		pFront, qFront := queue1[0], queue2[0]
		queue1, queue2 = queue1[1:], queue2[1:] // slices are great

		if pFront == nil && qFront == nil {
			continue
		}

		if pFront == nil || qFront == nil {
			return false
		}

		if pFront.Val != qFront.Val {
			return false
		}

		queue1 = append(queue1, pFront.Left, pFront.Right)
		queue2 = append(queue2, qFront.Left, qFront.Right)
	}

	return true
}

// Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/
func invertTree(root *TreeNode) *TreeNode {
	// approach one: traverse, swap left and right references for every node (recursive)
	// note: traversal (preorder, postorder, etc.) would be the same here
	// base case: if root == nil -> return nil
	// process root: swap left and right references, continue traversing those
	// no return value / stack frame build up value
	return invertTreeHelper(root)
}

func invertTreeHelper(root *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	// process node
	root.Left, root.Right = root.Right, root.Left

	invertTreeHelper(root.Left)
	invertTreeHelper(root.Right)

	return root
}

// Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/
func minWindow(s string, t string) string {
	// thinking...
	// minimum window -> makes me think we use a sliding window (2 ptrs) for this
	// including duplicates -> can't use a set or a hash on the characters really
	// order doesn't matter -> ABC can be BAC in the other string and that'd be the ans
	// can assume test case has a unique answer
	// can include both upper and lower case (those are considered different in the sol)

	// approach one: brute force = try every substring in s of at least len(t)
	// time = O(n^2) n substrings, each comparing characters
	// space = constant

	// approach 2: sliding window with a bunch of checks
	// create a substring window of string s; i = 0, j = 0
	// while we still have characters to read... (j < len(s))
	// - add char at j to running solution, update freq if exsts in map

	// - check if we have a working substring (charsLeft <= 0):
	// -- shrink the window (from start) while that remains true
	// -- remove freqs from map (but only if doing so wouldn't break our constraint)
	// -- then update sol if shorter

	// - if we have a solution -> increment both i and j, remove freq of s[i] from map and update charsLeft
	// - if we don't have a sol - only increment j
	// for each character (at most n) subtract from frequencies if in our map if exist

	// return sol

	// time complexity = O(m + n)
	// space = constant

	return minWindowApproachTwo(s, t)
}

func minWindowApproachTwo(s, t string) string {
	// edge case
	if len(t) > len(s) {
		return ""
	}

	// storage
	m, n := len(s), len(t)
	freqs := make(map[rune]int, n) // at most n unique letters

	// initialize frequencies to the ones we care about in t
	for _, r := range t {
		freqs[r] += 1
	}

	var sol string
	charsLeft := n

	// - charsLeft = len(t), each time we slide and check the letter, update if in map
	i, j := 0, 0
	for i <= j && j < m {
		c := rune(s[j])

		// if this character is also in t then count that frequency
		if freq, ok := freqs[c]; ok {
			freqs[c] = freq - 1

			// don't count it if it's a dup
			if freq > 0 {
				charsLeft--
			}
		}

		// note: charsLeft can be negative if we have multiple letters that are in t
		var first rune
		for charsLeft <= 0 {
			first = rune(s[i])

			// if our first character is irrelevant, just increment first
			if _, ok := freqs[first]; !ok {
				i++
				continue
			} else if freq, ok := freqs[first]; ok && freq < 0 {
				// if our first c is relevant, but duplicated, increment + update map
				i++
				freqs[first]++
			} else if freq, ok := freqs[first]; ok && freq == 0 {
				// else, capture our solution and move our start up
				if sol == "" {
					sol = s[i : j+1] // [inclusive idx, exlcusive)
				} else if j-i+1 < len(sol) {
					sol = s[i : j+1]
				}

				i++
				freqs[first]++
				charsLeft++
				break
			}
		}

		j++

	}

	return sol
}
