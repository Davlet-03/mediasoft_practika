package queue

// Queue реализует очередь по принципу FIFO:
// первым извлекается элемент, добавленный раньше остальных.
type Queue[T any] struct {
	items []T
}

// Enqueue добавляет элемент в конец очереди.
func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

// Dequeue удаляет и возвращает первый элемент.
// Второе значение равно false, если очередь пуста.
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}

	value := q.items[0]
	var empty T
	q.items[0] = empty // освобождаем ссылку на удалённый элемент
	q.items = q.items[1:]
	return value, true
}

// IsEmpty сообщает, пуста ли очередь.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Len возвращает количество элементов в очереди.
func (q *Queue[T]) Len() int {
	return len(q.items)
}
