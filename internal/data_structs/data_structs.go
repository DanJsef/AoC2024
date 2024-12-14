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

type Position struct {
	X int
	Y int
}

func (p Position) Add(p2 Position) Position {
	return Position{X: p.X + p2.X, Y: p.Y + p2.Y}
}

func (p Position) Sub(p2 Position) Position {
	return Position{X: p.X - p2.X, Y: p.Y - p2.Y}
}

func (p Position) AddWrap(p2 Position, width int, height int) Position {
	newPos := Position{X: (p.X + p2.X) % width, Y: (p.Y + p2.Y) % height}

	if newPos.X < 0 {
		newPos.X += width
	}

	if newPos.Y < 0 {
		newPos.Y += height
	}

	return newPos
}
