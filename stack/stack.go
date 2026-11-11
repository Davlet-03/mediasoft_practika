package stack

// Stack реализует стек по принципу LIFO:
// последний добавленный элемент извлекается первым.
type Stack[T any] struct {
	items []T
}

// Push добавляет элемент на вершину стека.
func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

// Pop удаляет и возвращает верхний элемент.
// Второе значение равно false, если стек пуст.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return value, true
}

// Peek возвращает верхний элемент, не удаляя его.
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty сообщает, пуст ли стек.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Len возвращает количество элементов в стеке.
func (s *Stack[T]) Len() int {
	return len(s.items)
}
