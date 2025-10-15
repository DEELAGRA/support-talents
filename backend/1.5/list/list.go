package main

import (
	"fmt"
)

type Node struct {
	Number int
	Next   *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) AddFirst(value int) {
	newNode := &Node{
		Number: value,
		Next:   ll.Head,
	}
	ll.Head = newNode
	fmt.Println("Добавил значение в начало")
}
func (ll *LinkedList) AddLast(value int) {

	newNode := &Node{
		Number: value,
		Next:   nil,
	}

	if ll.Head == nil {
		ll.Head = newNode
		fmt.Println("Добавил значение в конец")
		return
	}

	currnet := ll.Head
	for currnet.Next != nil {
		currnet = currnet.Next
	}
	currnet.Next = newNode

	fmt.Println("Добавил значение в конец")
}

func (ll *LinkedList) Print() {
	current := ll.Head

	for current != nil {
		fmt.Printf("%d -> ", current.Number)
		current = current.Next
	}

	fmt.Println("nil")
}
