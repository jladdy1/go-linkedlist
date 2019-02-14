package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello linked list")

	list := LinkedList{}
	list.Add(1)

	fmt.Printf("%v", list)
}
