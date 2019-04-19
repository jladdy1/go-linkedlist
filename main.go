package main

import (
	"fmt"
)

func main() {
	fmt.Println("Random Linked list stuff")

	list := LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(4)
	list.Add(10)
	list.Add(100)
	list.Add(5)
	list.Add(4)
	list.Add(45)

	//Just print the contents of the list
	printList(&list)

	kth, node := list.getKtolast(1)

	fmt.Printf("The kth from the last element is index %v and value %v", kth, node.data)

	printKtoLast(list.head, 2)

	//remove dups and print again

	fmt.Println("Removing Dups")
	removeDups2(&list)
	printList(&list)

}

func printList(list *LinkedList) {

	node := list.head
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

func removeDups(list *LinkedList) {
	//O(n) time complexity
	dic := make(map[int]*Node)

	node := list.head
	var previousNode *Node

	for node != nil {
		_, contains := dic[node.data]

		if contains {
			//remove
			previousNode.next = node.next
			node = previousNode.next
		} else {
			dic[node.data] = node
			previousNode = node
			node = node.next
		}
	}

}

func removeDups2(list *LinkedList) {
	// O(n^2) time complexity.  O(1) space complexity
	node := list.head

	for node != nil {
		runner := node

		for runner.next != nil {
			if runner.next.data == node.data {
				//dup
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}
		node = node.next
	}

}

func printKtoLast(head *Node, k int) int {

	if head == nil {
		return 0
	}

	index := printKtoLast(head.next, k) + 1

	if index == k {
		fmt.Printf("%v th to last is %v", k, head.data)
	}

	return index
}
