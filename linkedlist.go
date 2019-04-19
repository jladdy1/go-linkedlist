package main

//Node of a linked list
type Node struct {
	//This is an example of where generics would be nice in Go.
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

//Iterative solution
func (ll *LinkedList) getKtolast(k int) (int, *Node) {

	first, second := ll.head, ll.head
	for i := 0; i < k+1; i++ {
		second = second.next
		if second == nil {
			return 0, nil
		}
	}
	index := 1
	for second != nil {
		first = first.next
		second = second.next
		index++
	}
	return index, first
}
