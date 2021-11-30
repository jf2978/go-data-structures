package main

import (
	"fmt"
)

// PriorityNode is a "generic" element with an integer priority stored in a PriorityQueue
type PriorityNode struct {
	val      int
	priority int
}

// compareTo is priority comparator function that returns -1 if n < a, 0 if n == a and 1 if n > a
func (n *PriorityNode) compareTo(a *PriorityNode) int {
	if n.priority > a.priority {
		return 1
	} else if n.priority < a.priority {
		return -1
	} else {
		return 0
	}
}

// PriorityQueue is a heap implementation that stores PriorityNode values
type PriorityQueue struct {
	arr []*PriorityNode
}

// Insert adds an element to the heap
func (p *PriorityQueue) Insert(n *PriorityNode) {
	p.arr = append(p.arr, n)
	p.percolateUp(len(p.arr) - 1)
}

// ExtractMax removes the highest priority element from the queue
func (p *PriorityQueue) ExtractMax() *PriorityNode {
	max := p.arr[0]

	last := len(p.arr) - 1
	p.swap(0, last)

	// remember: slices are not arrays, but rather "flexible views into the elements of an array"
	// so here, we're re-assigning the heap slice to the same underlying data but excluding the last element
	p.arr = p.arr[:last]
	p.percolateDown(0)

	return max
}

// Search looks up the node with the provided value from the priority queue
// since heaps can't be traversed like BSTs (different constraints) this is just a linear lookup
func (p *PriorityQueue) Search(val int) *PriorityNode {
	for _, node := range p.arr {
		if node.val == val {
			return node
		}
	}

	return nil
}

// percolateUp recursively percolates up the priority queue
// swapping elements with its parent until parent >= element
func (h *PriorityQueue) percolateUp(i int) {
	// base case: if i is root, then we're done
	if i == 0 {
		return
	}

	parent := parent(i)

	// if the current element deserve to be the parent, swap em and continue percolating
	if h.arr[parent].compareTo(h.arr[i]) < 0 {
		h.swap(i, parent)
		h.percolateUp(parent)
	}
}

// percolateDown percolates down the max heap
// swapping the current element with its larger child until current >= both children
func (h *PriorityQueue) percolateDown(i int) {

	l, r := left(i), right(i)
	last := len(h.arr) - 1

	// while we have at least one left child
	var child int
	for l <= last {
		// pick the index of the larger child (or left if only one)
		if l == last {
			child = l
		} else if h.arr[l].compareTo(h.arr[r]) > 0 {
			child = l
		} else {
			child = r
		}

		// swap the current element with that child if it's smaller
		if h.arr[i].compareTo(h.arr[child]) < 0 {
			h.swap(i, child)
			i = child
			l, r = left(i), right(i)
		} else {
			return
		}
	}
}

// swap swaps elements i and j in the max heap
func (h *PriorityQueue) swap(i, j int) {
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

	// create a new max heap
	pq := PriorityQueue{}
	fmt.Printf("priority queue struct: %v\n", pq)

	// heapify some data
	data := []*PriorityNode{
		&PriorityNode{val: 100, priority: 20},
		&PriorityNode{val: 23, priority: 200},
		&PriorityNode{val: 1, priority: 10},
		&PriorityNode{val: 32, priority: 5},
		&PriorityNode{val: 56, priority: 12},
		&PriorityNode{val: 77, priority: 0},
	}
	for _, v := range data {
		pq.Insert(v.val) // ignore priority
		fmt.Printf("heap after inserting %v: %v\n", v, pq.arr)
	}

	// extract the min a few times
	pq.ExtractMax()
	pq.ExtractMax()
	pq.ExtractMax()
	pq.ExtractMax()
	pq.ExtractMax()
	pq.ExtractMax()
	pq.ExtractMax()
}
