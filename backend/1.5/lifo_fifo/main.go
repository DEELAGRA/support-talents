package main

import "fmt"

func main() {
	item := &Stack{}
	value := &Queue{}

	item.Push(31)
	item.Push(32)
	item.Push(33)
	item.Push(34)
	fmt.Println(item.Pop())

	value.Enqueue(61)
	value.Enqueue(62)
	value.Enqueue(63)
	value.Enqueue(64)
	fmt.Println(value.Dequeue())

}
