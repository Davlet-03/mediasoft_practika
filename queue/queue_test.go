package queue

import "testing"

func TestQueueEnqueueDequeue(t *testing.T) {
	var q Queue[int]
	q.Enqueue(10)
	q.Enqueue(20)

	if got, ok := q.Dequeue(); !ok || got != 10 {
		t.Fatalf("Dequeue() = %d, %v; want 10, true", got, ok)
	}
	if got, ok := q.Dequeue(); !ok || got != 20 {
		t.Fatalf("Dequeue() = %d, %v; want 20, true", got, ok)
	}
	if !q.IsEmpty() {
		t.Fatal("queue should be empty")
	}
}

func TestQueueWithStrings(t *testing.T) {
	var q Queue[string]
	q.Enqueue("first")
	if got, ok := q.Dequeue(); !ok || got != "first" {
		t.Fatalf("Dequeue() = %q, %v; want first, true", got, ok)
	}
}

func TestEmptyQueue(t *testing.T) {
	var q Queue[int]
	if _, ok := q.Dequeue(); ok {
		t.Fatal("Dequeue on empty queue should return false")
	}
}
