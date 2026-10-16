package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return value, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	var stack Stack

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Стек:", stack.items)

	value, ok := stack.Pop()
	if ok {
		fmt.Println("Удалили:", value)
	}

	value, ok = stack.Peek()
	if ok {
		fmt.Println("Верхний элемент:", value)
	}

	fmt.Println("Стек пустой:", stack.IsEmpty())
	fmt.Println()
	runQueueExample()
	fmt.Println()
	runLinkedListExample()
}
