package main

import (
	"fmt"
)

// MaxHeap is a max heap implementation that supports integer values
type MaxHeap struct {
	arr []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.percolateUp(len(h.arr) - 1)
}

// Extract removes the highest priority element from the max heap
func (h *MaxHeap) Extract() int {
	// remember: slices are not arrays, but rather "flexible views into the elements of an array"
	// so here, we're getting the root element, then re-assigning the heap array to the rest of the data
	var max int
	max, h.arr = h.arr[0], h.arr[1:]

	fmt.Printf("new slice: %v\n", h.arr)

	h.swap(0, len(h.arr)-1)
	h.percolateDown(0)

	return max
}

// percolateUp recursively percolates up the max heap
// swapping elements with its parent until parent >= element
func (h *MaxHeap) percolateUp(i int) {
	// base case: if i is root, then we're done
	if i == 0 {
		return
	}

	parent := parent(i)

	// if the current element deserve to be the parent, swap em and continue percolating
	if h.arr[parent] < h.arr[i] {
		h.swap(i, parent)
		h.percolateUp(parent)
	}
}

// percolateDown percolates down the max heap
// swapping the current element with its larger child until current >= both children
func (h *MaxHeap) percolateDown(i int) {

	l, r := left(i), right(i)
	last := len(h.arr) - 1

	// while we have at least one left child
	for l <= last {
		// pick the index of the larger child (or left if only one)
		var child int
		if l == last {
			child = l
		} else if h.arr[l] > h.arr[r] {
			child = l
		} else if h.arr[r] > h.arr[l] {
			child = r
		}

		// swap the current element with that child if it's smaller
		if h.arr[i] < h.arr[child] {
			h.swap(i, child)
			i = child
			l, r = left(i), right(i)
		} else {
			return
		}
	}
}

// swap swaps elements i and j in the max heap
func (h *MaxHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

// parent returns the index of the parent of the provided index i
func parent(i int) int {
	return (i - 1) / 2
}

// left returns the index of the left child of the provided index i
func left(i int) int {
	return 2 * i
}

// right returns the index of the right child of the provided index i
func right(i int) int {
	return (2 * i) + 1
}

func main() {

	// create a new max heap
	heap := MaxHeap{}

	// insert a bunch of stuff
	heap.Insert(100)
	heap.Insert(30)
	heap.Insert(205)
	heap.Insert(12)
	heap.Insert(23)
	heap.Insert(400)
	heap.Insert(150)
	heap.Insert(12)

	fmt.Printf("data: %+v\n", heap)

	// extract the max a few times
	heap.Extract()

	fmt.Printf("data: %+v\n", heap)

	heap.Extract()

	fmt.Printf("data: %+v\n", heap)

	heap.Extract()

	fmt.Printf("data: %+v", heap)
}
