package stack

import "testing"

func TestStackPushPop(t *testing.T) {
	var s Stack[int]
	s.Push(10)
	s.Push(20)

	if got, ok := s.Pop(); !ok || got != 20 {
		t.Fatalf("Pop() = %d, %v; want 20, true", got, ok)
	}
	if got, ok := s.Pop(); !ok || got != 10 {
		t.Fatalf("Pop() = %d, %v; want 10, true", got, ok)
	}
	if !s.IsEmpty() {
		t.Fatal("stack should be empty")
	}
}

func TestStackWithStrings(t *testing.T) {
	var s Stack[string]
	s.Push("Go")
	if got, ok := s.Peek(); !ok || got != "Go" {
		t.Fatalf("Peek() = %q, %v; want Go, true", got, ok)
	}
}

func TestEmptyStack(t *testing.T) {
	var s Stack[int]
	if _, ok := s.Pop(); ok {
		t.Fatal("Pop on empty stack should return false")
	}
	if _, ok := s.Peek(); ok {
		t.Fatal("Peek on empty stack should return false")
	}
}
