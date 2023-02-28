package stack

type Stack[T any] struct {
	stack []*T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{stack: make([]*T, 0)}
}

func (s *Stack[T]) Push(elem *T) {
	s.stack = append(s.stack, elem)
}

func (s *Stack[T]) Pop() *T {

	if s.IsNotEmpty() {
		last := len(s.stack) - 1
		elem := s.stack[last]
		s.stack = s.stack[:last]
		return elem
	}
	return s.stack[0]
}

func (s *Stack[T]) IsNotEmpty() bool {
	return len(s.stack) != 0
}

func (s *Stack[T]) Size() int {
	return len(s.stack)
}
