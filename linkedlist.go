package main

import (
	"container/list"
	"fmt"
)

// No need to reinvent the wheel by implementing a linked list in Go, but
// since I barely use the standard library list I figured it would still be a
// good exercise to use the data structure with some arbitrary data

// print prints all the Elements in the list starting from the front
func printList(l *list.List) {
	current := l.Front()
	fmt.Printf("(head) ")
	for current.Next() != nil {
		fmt.Printf("[%v] -> ", current.Value)
		current = current.Next()
	}
	fmt.Printf("nil\n")
}

// printReverse prints all the elements in the linked list in reverse order
func printListReverse(l *list.List) {
	current := l.Back()
	fmt.Printf("(tail) ")
	for current.Prev() != nil {
		fmt.Printf("[%v] -> ", current.Value)
		current = current.Prev()
	}
	fmt.Printf("nil\n")
}

func main() {
	ll := list.New()

	// push a single element (0) to mark the start
	ll.PushBack(0)

	// append 1-5 to the list
	for i := 1; i <= 5; i++ {
		ll.PushBack(i)
	}

	fmt.Printf("List Length: %d\n", ll.Len())

	// prepend negative values to the list
	for i := -1; i >= -5; i-- {
		ll.PushFront(i)
	}

	fmt.Printf("List Length: %d\n", ll.Len())

	printList(ll)
	printListReverse(ll)

	// remove a few elements in the list (need to remove by  reference to the Element)
	// arbitrarily remove every other element
	valuesToRemove := map[int]bool{
		4: true,
		2: true,
		0: true,
	}
	current := ll.Front()
	for current != nil {
		next := current.Next()
		if valuesToRemove[(current.Value).(int)] {
			fmt.Printf("found value: %d\n", current.Value)
			ll.Remove(current)
		}

		current = next
	}

	printList(ll)

}
