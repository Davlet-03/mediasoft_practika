package linkedlist

import (
	"reflect"
	"testing"
)

func TestAppendAndValues(t *testing.T) {
	var list LinkedList[int]
	list.Append(10)
	list.Append(20)
	list.Append(30)

	want := []int{10, 20, 30}
	if got := list.Values(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Values() = %v; want %v", got, want)
	}
	if list.Len() != 3 {
		t.Fatalf("Len() = %d; want 3", list.Len())
	}
}

func TestLinkedListWithStrings(t *testing.T) {
	var list LinkedList[string]
	list.Append("Go")
	list.Append("Generics")

	want := []string{"Go", "Generics"}
	if got := list.Values(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Values() = %v; want %v", got, want)
	}
}

func TestEmptyLinkedList(t *testing.T) {
	var list LinkedList[int]
	if !list.IsEmpty() {
		t.Fatal("new list should be empty")
	}
	if got := list.Values(); len(got) != 0 {
		t.Fatalf("Values() = %v; want empty slice", got)
	}
}
