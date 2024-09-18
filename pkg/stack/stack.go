package stack

type Stack[T any] struct {
	buffer []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	s.buffer = append(s.buffer, value)
}

func (s *Stack[T]) Pop() (result T, ok bool) {
	resultPtr, ok := s.Top()

	if !ok {
		return
	}

	result = *resultPtr
	s.buffer = s.buffer[:len(s.buffer)-1]
	return
}

func (s *Stack[T]) Top() (result *T, ok bool) {
	ok = len(s.buffer) != 0

	if !ok {
		return
	}

	result = &s.buffer[len(s.buffer)-1]
	return
}

func (s *Stack[T]) Len() int {
	return len(s.buffer)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}
