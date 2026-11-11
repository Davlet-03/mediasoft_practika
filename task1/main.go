package main

import (
	"fmt"

	"mediasoft_practice2/linkedlist"
	"mediasoft_practice2/queue"
	"mediasoft_practice2/stack"
)

func main() {
	var numbers stack.Stack[int]
	numbers.Push(10)
	numbers.Push(20)
	numbers.Push(30)

	fmt.Println("Стек:", numbers.Len(), "элементов")
	if value, ok := numbers.Pop(); ok {
		fmt.Println("Удалили из стека:", value)
	}
	if value, ok := numbers.Peek(); ok {
		fmt.Println("Верхний элемент стека:", value)
	}
	fmt.Println("Стек пустой:", numbers.IsEmpty())

	fmt.Println()

	var tasks queue.Queue[string]
	tasks.Enqueue("Первая задача")
	tasks.Enqueue("Вторая задача")
	fmt.Println("В очереди:", tasks.Len(), "элемента")
	if value, ok := tasks.Dequeue(); ok {
		fmt.Println("Взяли из очереди:", value)
	}
	fmt.Println("Очередь пустая:", tasks.IsEmpty())

	fmt.Println()

	var list linkedlist.LinkedList[int]
	list.Append(10)
	list.Append(20)
	list.Append(30)
	fmt.Println("Односвязный список:", list.Values())
	fmt.Println("Количество элементов:", list.Len())
}
