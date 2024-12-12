package utils

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

func (s *Stack[T]) Add(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
