package main

import "fmt"

// MaxHeap is a max heap implementation that supports integer values
type MaxHeap struct {
	arr []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.percolateUp(len(h.arr) - 1)
}

// percolateUp recursively percolates up the max heap
// swapping elements with its parent until parent >= element
func (h *MaxHeap) percolateUp(i int) {
	// base case: if i is root, then we're done
	if i == 0 {
		return
	}

	parent := parent(i)

	// base case: if our parent isn't in the right place, swap and keep going
	if h.arr[parent] < h.arr[i] {
		h.swap(i, parent)
		h.percolateUp(parent)
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
	heap.Insert(100)
	heap.Insert(30)
	heap.Insert(205)
	heap.Insert(12)
	heap.Insert(23)

	fmt.Printf("data: %+v", heap)

}
