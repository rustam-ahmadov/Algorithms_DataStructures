package main

import "fmt"

type node struct {
	v    int
	next *node
}

type list struct {
	head *node
	tail *node
	size int
}

func (list *list) add(v int) {
	newNode := &node{v, nil}

	if list.head == nil {
		list.head = newNode
		list.tail = list.head
		return
	}

	temp := list.tail
	temp.next = newNode
	list.tail = temp.next
}

func (list *list) print() {
	for temp := list.head; temp != nil; temp = temp.next {
		fmt.Printf("%d\n", temp.v)
	}
}

func main() {
	list := new(list)
	list.add(1)
	list.add(2)
	list.add(3)
	list.print()
}
