package utils

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) UnionUpdate(t Set[T]) {
	for v := range t {
		s[v] = struct{}{}
	}
}
