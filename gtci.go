package main

import "fmt"

// practice coding questions from Grokking the Coding Interview
// some of these are repeats of what is on LeetCode, but repitition of the same pattern
// over and over again will hammer in my understanding of the basic problem being solved!

/** SLIDING WINDOW */

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

func main() {
	k1 := 3
	arr1 := []int{2, 1, 5, 1, 3, 2}

	ans1 := findMaxSubArray(k1, arr1) // expected output: 9
	fmt.Printf("findMaxSubArray(%d, %v) = %d\n", k1, arr1, ans1)

	k2 := 2
	arr2 := []int{2, 3, 4, 1, 5}

	ans2 := findMaxSubArray(k2, arr2) // expected output: 7
	fmt.Printf("findMaxSubArray(%d, %v) = %d\n", k2, arr2, ans2)

}
