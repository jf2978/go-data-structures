package main

import (
	"fmt"
)

// MinHeap is a min heap implementation that supports integer values
type MinHeap struct {
	arr []int
}

// Insert adds an element to the heap
func (h *MinHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.percolateUp(len(h.arr) - 1)

	fmt.Printf("heap after inserting %v: %v\n", val, h.arr)
}

// Extract removes the highest priority element from the max heap
func (h *MinHeap) Extract() int {

	min := h.arr[0]
	fmt.Printf("extracting min: %v\n", min)

	last := len(h.arr) - 1
	h.swap(0, last)

	// remember: slices are not arrays, but rather "flexible views into the elements of an array"
	// so here, we're re-assigning the heap slice to the same underlying data but excluding the last element
	h.arr = h.arr[:last]
	h.percolateDown(0)

	fmt.Printf("new heap: %v\n", h.arr)

	return min
}

// percolateUp recursively percolates up the min heap
// swapping elements with its parent until parent >= element
func (h *MinHeap) percolateUp(i int) {
	// base case: if i is root, then we're done
	if i == 0 {
		return
	}

	parent := parent(i)

	// if the current element deserves to be the parent instead, swap and continue percolating
	if h.arr[i] < h.arr[parent] {
		h.swap(i, parent)
		h.percolateUp(parent)
	}
}

// percolateDown percolates down the min heap
// swapping the current element with its larger child until current >= both children
func (h *MinHeap) percolateDown(i int) {

	l, r := left(i), right(i)
	last := len(h.arr) - 1

	// while we have at least one left child
	for l <= last {
		// pick the index of the smaller child (or left if only one)
		var child int
		if l == last {
			child = l
		} else if h.arr[l] < h.arr[r] {
			child = l
		} else if h.arr[r] < h.arr[l] {
			child = r
		}

		// swap the current element with that child if it's larger
		if h.arr[i] > h.arr[child] {
			h.swap(i, child)
			i = child
			l, r = left(i), right(i)
		} else {
			return
		}
	}
}

// swap swaps elements i and j in the min heap
func (h *MinHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

// parent returns the index of the parent of the provided index i
func parent(i int) int {
	return (i - 1) / 2
}

// left returns the index of the left child of the provided index i
func left(i int) int {
	return (2 * i) + 1
}

// right returns the index of the right child of the provided index i
func right(i int) int {
	return (2 * i) + 2
}

func main() {

	// create a new min heap
	heap := MinHeap{}

	// heapify some data
	data := []int{100, 30, 205, 12, 23, 400, 150, 12}
	for _, v := range data {
		heap.Insert(v)
	}

	fmt.Printf("data: %+v\n", heap)

	// extract the min a few times
	heap.Extract()
	heap.Extract()
	heap.Extract()
	heap.Extract()
	heap.Extract()
	heap.Extract()
	heap.Extract()
	heap.Extract()
}
