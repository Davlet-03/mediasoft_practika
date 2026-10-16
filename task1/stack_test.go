package main

import "testing"

func TestStackPushPop(t *testing.T) {
	var stack Stack

	stack.Push(10)
	stack.Push(20)

	value, ok := stack.Pop()
	if !ok || value != 20 {
		t.Fatalf("Pop() = %d, %v; ожидали 20, true", value, ok)
	}

	value, ok = stack.Pop()
	if !ok || value != 10 {
		t.Fatalf("Pop() = %d, %v; ожидали 10, true", value, ok)
	}

	if !stack.IsEmpty() {
		t.Fatal("стек должен быть пустым")
	}
}

func TestEmptyStack(t *testing.T) {
	var stack Stack

	if _, ok := stack.Pop(); ok {
		t.Fatal("Pop() пустого стека должен вернуть false")
	}

	if _, ok := stack.Peek(); ok {
		t.Fatal("Peek() пустого стека должен вернуть false")
	}
}
