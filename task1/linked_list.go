package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{value: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (l *LinkedList) Print() {
	current := l.head

	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}

	fmt.Println("nil")
}

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

func runLinkedListExample() {
	var list LinkedList

	list.Append(10)
	list.Append(20)
	list.Append(30)

	fmt.Print("Односвязный список: ")
	list.Print()

	fmt.Println("Список пустой:", list.IsEmpty())
}
