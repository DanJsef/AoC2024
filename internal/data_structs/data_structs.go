package datastructs

type node[T any] struct {
	next  *node[T]
	value T
}

type Stack[T any] struct {
	start *node[T]
	count int
}

func (s *Stack[T]) Push(value T) {
	newNode := &node[T]{value: value, next: s.start}
	s.start = newNode
	s.count++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.count == 0 {
		var empty T
		return empty, false
	}
	value := s.start.value
	s.start = s.start.next
	s.count--
	return value, true
}

func (s *Stack[T]) Peak() (T, bool) {
	if s.count == 0 {
		var empty T
		return empty, false
	}
	return s.start.value, true
}

func (s *Stack[T]) Len() int {
	return s.count
}
