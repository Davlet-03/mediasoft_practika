package linkedlist

// Node — узел односвязного списка.
type Node[T any] struct {
	value T
	next  *Node[T]
}

// LinkedList — односвязный список элементов типа T.
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

// Append добавляет элемент в конец списка за O(1).
func (l *LinkedList[T]) Append(value T) {
	node := &Node[T]{value: value}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.len++
}

// Values возвращает элементы списка в виде среза.
func (l *LinkedList[T]) Values() []T {
	values := make([]T, 0, l.len)
	for current := l.head; current != nil; current = current.next {
		values = append(values, current.value)
	}
	return values
}

// IsEmpty сообщает, пуст ли список.
func (l *LinkedList[T]) IsEmpty() bool {
	return l.len == 0
}

// Len возвращает количество элементов в списке.
func (l *LinkedList[T]) Len() int {
	return l.len
}
