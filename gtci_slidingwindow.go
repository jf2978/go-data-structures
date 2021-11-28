package main

import (
	"fmt"
)

// practice coding questions from Grokking the Coding Interview
// some of these are repeats of what is on LeetCode, but repitition of the same pattern
// over and over again will hammer in my understanding of the basic problem being solved!

// Given an array of positive numbers and a positive number ‘k,’
// find the maximum sum of any contiguous subarray of size ‘k’.
func findMaxSubArray(k int, arr []int) int {
	// approach: edge case 1: if k > len(arr) return -1
	// loop from 0 to k - 1 and get the starting sum
	// loop for i = 1, j = k, range through arr i++ j++,
	// subtract arr[i-1] from running sum
	// add arr[j] to running sum
	// if > currentMax, update the value
	// return currentMax

	if k > len(arr) {
		return -1
	}

	// compute starting sum
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	max := sum

	// slide window and check against maxiumum sum
	for i, j := 1, k; j < len(arr); i, j = i+1, j+1 {
		sum -= arr[i-1]
		sum += arr[j]

		if sum > max {
			max = sum
		}
	}

	return max
}

// Given a string, find the length of the longest substring in it with no more than K distinct characters.
// derivative of leetcode question: lengthOfLongestSubstringOne (already solved)
func longestSubstringKDistinct(s string, k int) int {
	// approach one: store distinct characters as a map char -> list of indices
	// iterate through the string (one rune at a time):
	// add current character to the map char -> append(current j index) to tail of linked list

	// while set size > k:
	// lookup char in map and remove from front pf the list; if list is empty after this, delete the entry entirely from the map
	// increment i

	// in any case, update maxLength with j - i + 1 if it's > maxLength
	// (implied) if set size <= k then just keep expanding the window (j++), we allow eq here because repeat characters could extend this window

	maxLength := 0
	mem := make(map[rune][]int, len(s)) // maps char -> indices where this character occurs
	for i, j := 0, 0; j < len(s); j++ {
		lastChar := rune(s[j])
		mem[lastChar] = append(mem[lastChar], j)

		// if we have too many distinct characters, slide the window + update our storage
		for len(mem) > k {
			firstChar := rune(s[i])
			indices := mem[firstChar]

			indices = indices[1:]
			if len(indices) == 0 {
				delete(mem, firstChar) // remove entry from map entirely if there are no more occurrences
			}

			i++
		}

		if j-i+1 > maxLength {
			maxLength = j - i + 1
		}
	}

	return maxLength

	// note: better way to do this would've been to just *count* the occurrences in the map instead of having a list of all the indices
	// this way was a little more convoluted, but it worked!
}

// Given an array of positive numbers and a positive number ‘S,’ find the length
// of the smallest contiguous subarray whose sum is greater than or equal to ‘S’.
// Return 0 if no such subarray exists.
func findMinSubarray(s int, arr []int) int {
	// approach: loop through arr, storing running sum and min
	// add to the sum and increment the end index until sum >= s
	// once that happens try shrinking the window until we break the condition, update min if it gets better

	sum, min := 0, len(arr)+1
	for i, j := 0, 0; j < len(arr); j++ {
		sum += arr[j]

		// if our sum meets the threshold, update length if it's better and shrink the window
		for sum >= s {
			l := j - i + 1
			if l < min {
				min = l
			}

			sum -= arr[i]
			i++
		}
	}

	// if we've iterated through and min hasn't been updated, then no subarray exists
	if min > len(arr) {
		return 0
	}

	return min
}

// Given an array of characters where each character represents a fruit tree,
// you are given two baskets, and your goal is to put maximum number of fruits in each basket.
// The only restriction is that each basket can have only one type of fruit.

// You can start with any tree, but you can’t skip a tree once you have started. You will pick one fruit from each tree until you cannot,
// i.e., you will stop when you have to pick from a third fruit type.
func maxFruit(fruits []rune) int {
	// approach one: try every combination of fruit combos (basically like substrings)
	// time = O(n^2) which is bad

	// approach two: this is basically asking "whats the longest substring with at most 2 unique characters"
	// same approach as longestSubstringKDistinct except constrained to where K = 2 and
	// handling it as an array of runes directly instead of a string (implicitly the same thing)
	// I''ll re-implement this using the better approach I found after the fact (using a map that counts occurrences)

	maxFruits := 0
	mem := make(map[rune]int, len(fruits)) // maps char -> # of occurrences of this character
	for i, j := 0, 0; j < len(fruits); j++ {
		lastFruit := fruits[j]
		mem[lastFruit] = mem[lastFruit] + 1

		// if we're already at our 2 fruit limit, shrink the window + update our storage
		for len(mem) > 2 {
			firstChar := fruits[i]
			mem[firstChar] = mem[firstChar] - 1

			if mem[firstChar] == 0 {
				delete(mem, firstChar)
			}

			i++
		}

		if j-i+1 > maxFruits {
			maxFruits = j - i + 1
		}
	}

	return maxFruits
}

func main() {
	k1 := 3
	arr1 := []int{2, 1, 5, 1, 3, 2}

	ans1 := findMaxSubArray(k1, arr1) // expected output: 9
	fmt.Printf("findMaxSubArray(%d, %v) = %d\n", k1, arr1, ans1)

	k2 := 2
	arr2 := []int{2, 3, 4, 1, 5}

	ans2 := findMaxSubArray(k2, arr2) // expected output: 7
	fmt.Printf("findMaxSubArray(%d, %v) = %d\n", k2, arr2, ans2)

	s1 := 7
	arr3 := []int{2, 1, 5, 2, 3, 2}

	ans3 := findMinSubarray(s1, arr3) // expected output: 2
	fmt.Printf("findMinSubarray(%d, %v) = %d\n", s1, arr3, ans3)

	arr4 := []int{2, 1, 5, 2, 8}

	ans4 := findMinSubarray(s1, arr4) // expected output: 1
	fmt.Printf("findMinSubarray(%d, %v) = %d\n", s1, arr4, ans4)

	s3 := 8
	arr5 := []int{3, 4, 1, 1, 6}

	ans5 := findMinSubarray(s3, arr5) // expected output: 3
	fmt.Printf("findMinSubarray(%d, %v) = %d\n", s3, arr5, ans5)

	str1 := "araaci"
	ans6 := longestSubstringKDistinct(str1, k2)
	fmt.Printf("longestSubstringKDistinct(%s, %d) = %d\n", str1, k2, ans6)

	k3 := 1
	ans7 := longestSubstringKDistinct(str1, k3)
	fmt.Printf("longestSubstringKDistinct(%s, %d) = %d\n", str1, k3, ans7)

	str2 := "cbbebi"
	k4 := 3
	ans8 := longestSubstringKDistinct(str2, k4)
	fmt.Printf("longestSubstringKDistinct(%s, %d) = %d\n", str2, k4, ans8)

	k5 := 10
	ans9 := longestSubstringKDistinct(str2, k5)
	fmt.Printf("longestSubstringKDistinct(%s, %d) = %d\n", str2, k5, ans9)

	fruits1 := []rune{'A', 'B', 'C', 'A', 'C'}
	ans10 := maxFruit(fruits1)
	fmt.Printf("maxFruit(%s) = %d\n", string(fruits1), ans10)

	fruits2 := []rune{'A', 'B', 'C', 'B', 'B', 'C'}
	ans11 := maxFruit(fruits2)
	fmt.Printf("maxFruit(%s) = %d\n", string(fruits2), ans11)

}
