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

	s2 := 7
	arr4 := []int{2, 1, 5, 2, 8}

	ans4 := findMinSubarray(s2, arr4) // expected output: 1
	fmt.Printf("findMinSubarray(%d, %v) = %d\n", s2, arr4, ans4)

	s3 := 8
	arr5 := []int{3, 4, 1, 1, 6}

	ans5 := findMinSubarray(s3, arr5) // expected output: 3
	fmt.Printf("findMinSubarray(%d, %v) = %d\n", s3, arr5, ans5)
}
