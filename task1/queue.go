package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}

	value := q.items[0]
	q.items = q.items[1:]

	return value, true
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func runQueueExample() {
	var queue Queue

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Очередь:", queue.items)

	value, ok := queue.Dequeue()
	if ok {
		fmt.Println("Удалили из очереди:", value)
	}

	fmt.Println("Очередь после удаления:", queue.items)
	fmt.Println("Очередь пустая:", queue.IsEmpty())
}
