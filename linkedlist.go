package main

//Node of a linked list
type Node struct {
	next *Node
	data int
}

//LinkedList - simple linked list
type LinkedList struct {
	length int
	head   *Node
}

//Add a value to the list
func (ll *LinkedList) Add(d int) {
	n := Node{data: d}
	if ll.length == 0 {
		ll.head = &n
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = &n
	}
	ll.length++
}
